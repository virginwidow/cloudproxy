//
//  accessControl.cpp
//      John Manferdelli
//
//  Description: Access control implementation
//
//  Copyright (c) 2011, Intel Corporation. Some contributions 
//    (c) John Manferdelli.  All rights reserved.
//
// Use, duplication and disclosure of this file and derived works of
// this file are subject to and licensed under the Apache License dated
// January, 2004, (the "License").  This License is contained in the
// top level directory originally provided with the CloudProxy Project.
// Your right to use or distribute this file, or derived works thereof,
// is subject to your being bound by those terms and your use indicates
// consent to those terms.
//
// If you distribute this file (or portions derived therefrom), you must
// include License in or with the file and, in the event you do not include
// the entire License in the file, the file must contain a reference
// to the location of the License.


#include "jlmTypes.h"
#include "jlmUtility.h"
#include "logging.h"
#include "jlmcrypto.h"
#include "cert.h"
#include "resource.h"
#include "request.h"
#include "signedAssertion.h"
#include "accessControl.h"
#include "vault.h"
#include "cryptoHelper.h"
#include "validateEvidence.h"

#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <time.h>


// ----------------------------------------------------------------------------


/*
 *  Verbs: read | write | own | delete | create | request | has | is-a | get-keys | get-metadata
 *  
 *  Principal says Principal verb resource if condition
 *  Principal says Principal may say other-Principal verb object condition
 *  Example
 *      K-Policy says K-JLM can read //www.manferdelli.com/Files
 *      K-Policy says K-JLM can write //www.manferdelli.com/Files
 *      K-Policy says K-JLM can say x owns //www.manferdelli.com/Files
 *  <Assertions count='5'>
 *    <Assertion> K-Policy says K-JLM may read //www.manferdelli.com/Files </Assertion>
 *    <Assertion> K-Policy says K-JLM may read //www.manferdelli.com/Files </Assertion>
 *    <Assertion> K-Policy says K-JLM may read //www.manferdelli.com/Files/\* </Assertion>
 *    <Assertion> K-Policy says K-JLM may write //www.manferdelli.com/Files </Assertion>
 *    <Assertion> K-Policy says K-JLM may say x owns //www.manferdelli.com/Files </Assertion>
 *  </Assertions>
 *
 *  Request
 *    read //www.manferdelli.com/Files/MachineName/file
 */


inline bool whitespace(char b)
{
    return(b==' ' || b=='\t' || b=='\r' || b=='\n');
}


int nextToken(const char* sz, const char** pszToken)
{
    int     n;

    if(sz==NULL)
        return -1;

    while(*sz!='\0') {
        if(!whitespace(*sz))
            break;
    sz++;
    }

    if(*sz=='\0')
        return -1;

    *pszToken= sz;
    n= 0;
    while(*sz!='\0' && !whitespace(*sz)) {
        sz++;
        n++;
    }
    return n;
}


char* verbName(u32 uVerb)
{
    switch(uVerb) {
      case MAYREAD:
        return (char*)"mayread";
      case MAYWRITE:
        return (char*)"maywrite";
      case MAYCREATE:
        return (char*)"maycreate";
      case MAYDELETE:
        return (char*)"maydelete";
      case MAYOWN:
        return (char*)"mayown";
      case SPEAKSFOR:
        return (char*)"speaksfor";
      case SAYS:
        return (char*)"says";
      default:
        return (char*)"unknown verb";
    }
}


u32 verbFlag(const char* pVerbName)
{
   if(strcmp(pVerbName,"mayread")==0)
        return MAYREAD;

   if(strcmp(pVerbName, "maywrite")==0)
        return MAYWRITE;

   if(strcmp(pVerbName, "maycreate")==0)
        return MAYCREATE;

   if(strcmp(pVerbName, "maydelete")==0)
        return MAYDELETE;

   if(strcmp(pVerbName, "mayown")==0)
        return MAYOWN;

   if(strcmp(pVerbName, "speaksfor")==0)
        return SPEAKSFOR;

   if(strcmp(pVerbName, "says")==0)
        return SAYS;

    return 0;
}



// ---------------------------------------------------------------------------


assertionNode::assertionNode()
{
    m_fValidated= false;
    m_pAssertion= NULL;
    m_uVerbs= 0;
    m_pPrincipal= NULL;
    m_pResource= NULL;
    m_szCondition= NULL;
}


assertionNode::~assertionNode()
{
}


bool assertionNode::parseAssertion(accessPrincipal* pPrincipalSays, 
                    const char* szAssertion, bool fValidated)
{
    char                szBuf[MAXTOKEN];
    const char*         szTok= NULL;
    accessPrincipal*    pPrinc= NULL;
    resource*           pResource= NULL;
    int                 n;

#ifdef RULESTEST1
    fprintf(g_logFile, "parseAssertion %s\n", szAssertion);
    if(pPrincipalSays!=NULL)
        fprintf(g_logFile, "%s says\n", pPrincipalSays->getName());
#endif

    m_fValidated= fValidated;
    // who says
    if(pPrincipalSays!=NULL) {
        m_pPrincipal= pPrincipalSays;
        m_uVerbs= SAYS;
        m_pAssertion= new assertionNode();
        return m_pAssertion->parseAssertion((accessPrincipal*)NULL, szAssertion, 
                    (bool)fValidated);
    }

    // Subject
    n= nextToken(szAssertion, &szTok);
    if(n<0 || n>=MAXTOKEN)
        return false;
    memcpy(szBuf, szTok, n); szBuf[n]= '\0';
#if 0
    pPrinc= m_pMeta->findPrincipal(szBuf);
#endif
    // jlm READTHIS: should this really return true? That means the assertion was
    // parsed correctly but the resource in the assertion is NULL. I would think
    // that any syntactically well-formed assertion should parse and the resource
    // should be created. Ownership can be determined by the semantics of the 
    // assertion
    if(pPrinc==NULL) {
        fprintf(g_logFile, "No subject in assertion\n");
        return true;
    }
    m_pPrincipal= pPrinc;

    szAssertion= szTok+n;

    // Verb
    n= nextToken(szAssertion, &szTok);
    if(n<0 || n>=MAXTOKEN)
        return false;
    memcpy(szBuf, szTok,n); szBuf[n]= '\0';
    m_uVerbs= verbFlag(szBuf);

    szAssertion= szTok+n;
    if(strcmp(szBuf, "says")==0 || strcmp(szBuf, "maysay")==0) {
        // Fix: what is supposed to happen here? Is this meant to be a check 
        // for things that aren't 'says' or 'maysay'?
    }

    // Resource
    n= nextToken(szAssertion, &szTok);
    if(n<0 || n>=MAXTOKEN)
        return false;
    memcpy(szBuf, szTok,n); szBuf[n]= '\0';
#if 0
    pResource= m_pMeta->findResource(szBuf);
#endif
    if(pResource==NULL) {
        fprintf(g_logFile, "Could not find resource %s, so not a valid assertion\n", szBuf);
        return true;
    }
    m_pResource= pResource;

    // Fix: Add condition
    return true;
}


#ifdef TEST
void    assertionNode::printMe()
{
    if(m_fValidated)
        fprintf(g_logFile, "Assertion is validated\n");
    else
        fprintf(g_logFile, "Assertion is not validated\n");

    if(m_pPrincipal!=NULL)
        fprintf(g_logFile, "%s ", m_pPrincipal->m_szPrincipalName);

    fprintf(g_logFile, "%s ", verbName(m_uVerbs));
    if(m_pAssertion!=NULL) {
        fprintf(g_logFile, "\n");
        m_pAssertion->printMe();
        return;
    }

    if(m_pResource!=NULL)
        fprintf(g_logFile, "%s\n", m_pResource->m_szResourceName);
}
#endif


bool assertionNode::matchAction(u32 uVerb)
{
    if((uVerb&m_uVerbs)!=0)
        return true;
    return false;
}


bool assertionNode::matchResource(resource* pResource)
{
    if (NULL == pResource || NULL == m_pResource) return false;
    if(strcmp(pResource->m_szResourceName, m_pResource->m_szResourceName)==0)
        return true;
    return false;
}


bool assertionNode::matchPrincipal(accessPrincipal* pSubject)
{
    if (NULL == pSubject || NULL == m_pPrincipal) return false;
    if(pSubject->m_fValidated && m_pPrincipal->m_fValidated &&
       strcmp(pSubject->m_szPrincipalName, m_pPrincipal->m_szPrincipalName)==0)
        return true;
    return false;
}


bool assertionNode::assertionSucceeds(accessPrincipal* pSubject, u32 uVerb, resource* pResource,
                            int iNumAssertions, assertionNode** rgpAssertions)
{
#ifdef RULESTEST
    fprintf(g_logFile, "accessSucceeds %d\n", iNumAssertions);
#endif
    if(!m_fValidated) {
        fprintf(g_logFile, "Assertion not validated\n");
        return false;
        }

    // K1 says K2 speaksfor verb resource condition
    // Fix: speaksfor
    if((m_uVerbs&SPEAKSFOR)!=0) {
    }

    // K1 says K2 verb resource condition
    if((m_uVerbs&SAYS)!=0) {
        if(m_pAssertion==NULL) {
            fprintf(g_logFile, "empty subassertion\n");
            return false;
        }
        if(!m_pAssertion->assertionSucceeds(pSubject, uVerb, pResource, iNumAssertions, rgpAssertions)) {
            return false;
}
        if(isAnOwner(m_pPrincipal, pResource) || 
           isPolicyPrincipal(m_pPrincipal))
            return true;
        if(iNumAssertions>0)
            return m_pAssertion->assertionSucceeds(m_pPrincipal, uVerb, pResource, 
                                                   iNumAssertions-1, &rgpAssertions[1]);
        return false;
    }

    // K verb resource
    if(!matchAction(uVerb)) {
        fprintf(g_logFile, "verb mismatch %04x %04x\n", uVerb, m_uVerbs);
        return false;
    }
    if(!matchPrincipal(pSubject)) {
        fprintf(g_logFile, "principal mismatch\n");
        return false;
    }

    fprintf(g_logFile, "About to try and check pResource %p against m_pResource %p\n", pResource, m_pResource);
    fflush(g_logFile);     
 
    if (!matchResource(pResource)) {
        fprintf(g_logFile, "resource mismatch\n");
        return false;
    }

#ifdef RULESTEST
    fprintf(g_logFile, "assertionSucceeds returns true\n");
#endif
    return true;
}


bool isAnOwner(accessPrincipal* pSubject, resource* pResource)
{
    aNode<accessPrincipal>* pOwnerNode= pResource->m_myOwners.pFirst;
    accessPrincipal*        pOwnerPrincipal= NULL;

    while(pOwnerNode!=NULL) {
        pOwnerPrincipal= pOwnerNode->pElement;
#ifdef RULESTEST1
        fprintf(g_logFile, "Owner\n");
        pOwnerPrincipal->printMe();
#endif
        if(pSubject->m_fValidated && pOwnerPrincipal->m_fValidated &&
           strcmp(pSubject->m_szPrincipalName, pOwnerPrincipal->m_szPrincipalName)==0 &&
           pSubject->m_uPrincipalType==pOwnerPrincipal->m_uPrincipalType) {
            return true;
        }
        pOwnerNode= pOwnerNode->pNext;
    }
    return false;
}


bool isPolicyPrincipal(accessPrincipal* pSubject)
{
#if 0
    if(strcmp(pSubject->m_szPrincipalName, g_policyAccessPrincipal->m_szPrincipalName)==0)
        return true;
#endif
    return false;
}


// ---------------------------------------------------------------------------


accessRequest::accessRequest()
{
    m_szSubject= NULL;
    m_szRequest= NULL;
    m_szResource= NULL;
}


accessRequest::~accessRequest()
{
    if(m_szSubject==NULL) {
        free(m_szSubject);
        m_szSubject= NULL;
    }
    if(m_szResource==NULL) {
        free(m_szResource);
        m_szResource= NULL;
    }
}


#ifdef TEST
void accessRequest::printMe()
{
    fprintf(g_logFile, "\n\taccessRequest\n");
    if(m_szSubject==NULL)
        fprintf(g_logFile, "\tSubject is NULL\n");
    else
        fprintf(g_logFile, "\tSubject is %s\n", m_szSubject);
    if(m_szRequest==NULL)
        fprintf(g_logFile, "\tRequest is NULL\n");
    else
        fprintf(g_logFile, "\tRequest is %s\n", m_szRequest);
    if(m_szResource==NULL)
        fprintf(g_logFile, "\tResource is NULL\n");
    else
        fprintf(g_logFile, "\tResource is %s\n", m_szResource);
    fprintf(g_logFile, "\n");
}
#endif


accessGuard::accessGuard()
{
    m_fValid= false;
    m_pMeta= NULL;
    m_iNumAssertions= 0;
    m_iNumSubjects= 0; 
    m_rgpAssertions= NULL;
}


accessGuard::~accessGuard()
{
}


bool accessGuard::initChannelGuard(int numPrin, PrincipalCert** rgPrincs,
                                   RSAKey* pPolicy, metaData* pMeta)
{
    int                 iNumSubj;
    PrincipalCert**     rgpPrinc;
    accessPrincipal*    pPrincipal= NULL;
    PrincipalCert*      pCert= NULL;
    int                 i;

#ifdef TEST  
    fprintf(g_logFile, "initChannelGuard\n");
    fflush(g_logFile);
#endif

#if 0
    if(m_pMeta==NULL || m_pSession) {
        fprintf(g_logFile, "initChannelGuard: missing metaData or session\n");
        return false;
    }

    // register principals
    if(m_pserverCert!=NULL) {
        if(registerPrincipalfromCert(m_pserverCert)==NULL)
            throw "session::clientprotocolNego: Can't register server principal\n";
    }

    if(registerPrincipalfromCert(m_pclientCert)==NULL)
        throw "session::clientprotocolNego: Can't register client principal\n";

    for(i=0;i<m_iNumPrincipals; i++) {
        if(m_rgPrincipalCerts[i]!=NULL) {
            if(registerPrincipalfromCert(m_rgPrincipalCerts[i])==NULL)
                throw "session::clientprotocolNego: Can't register client principal\n";
        }
    }
#endif

    m_iNumSubjects= iNumSubj;
    for(i=0;i< iNumSubj; i++) {
        pCert= rgpPrinc[i];
        if(pCert==NULL) {
            fprintf(g_logFile, "initChannelGuard: NULL principal\n");
            return false;
        }
        if((pPrincipal= m_pMeta->findPrincipal(pCert->getPrincipalName()))==NULL) {
            pPrincipal= principalFromCert(pCert, true);
            if(!m_pMeta->addPrincipal(pPrincipal)) {
                fprintf(g_logFile, "initChannelGuard: can't add principal\n");
                return false;
            }
            if(pPrincipal==NULL) {
                fprintf(g_logFile, "initChannelGuard: pPrincipal is NULL\n");
                return false;
            }
        }
        m_Subjects.append(pPrincipal);
    }

    m_fValid= true;
    return true;
}


bool  accessGuard::permitAccess(accessRequest& req, const char* szCollection)
{
    resource*               pResource= NULL;
    aNode<accessPrincipal>* pSubjNode= NULL;
    aNode<accessPrincipal>* pOwnerNode= NULL;
    accessPrincipal*        pSubjPrincipal= NULL;
    accessPrincipal*        pSaysPrincipal= NULL;
    int                     i;
    u32                     uVerb= 0;

#ifdef TEST
    fprintf(g_logFile, "permitAccess\n");
    req.printMe();
    fprintf(g_logFile, "szCollection: %s\n", szCollection);
#endif

    if(!m_fValid) {
        fprintf(g_logFile, "permitAccess: accessGuard invalid\n");
        return false;
    }

    // accessPrincipals should have been validated by now
    pResource= m_pMeta->findResource(req.m_szResource);
    if(pResource==NULL) {
        fprintf(g_logFile, "permitAccess resource is NULL\n");
        return false;
    }
    pOwnerNode= pResource->m_myOwners.pFirst;
    UNUSEDVAR(pOwnerNode);
    pSubjNode= m_Subjects.pFirst;

    // Owners get all rights, we don't need no stinking evidence
    while(pSubjNode!=NULL) {
        pSubjPrincipal= pSubjNode->pElement;
        if(isAnOwner(pSubjPrincipal, pResource)) {
            fprintf(g_logFile, 
                    "permitAccess: The subject %s is an owner of resource %s, so the access check passes\n", 
                    pSubjPrincipal->m_szPrincipalName, req.m_szResource);
            return true;
        }       
        pSubjNode= pSubjNode->pNext;
    }

    // if request is add or delete owner, return false
    //      only owners have this right.
    if(strcmp(req.m_szRequest, "addOwner")==0 || 
       strcmp(req.m_szRequest, "removeOwner")==0) {
        fprintf(g_logFile, "permitAccess: no Evidence\n");
        return false;
    }

    // Does evidence support access?
    if(szCollection==NULL) {
        fprintf(g_logFile, "permitAccess: no Evidence\n");
        return false;
    }

    // parse evidence
    SignedAssertion*    pAssert= NULL; 
    evidenceCollection  oEvidenceCollection;

    if(!oEvidenceCollection.parseEvidenceCollection(szCollection)) {
        fprintf(g_logFile, "permitAccess: Can't parse Evidence list\n");
        return false;
    }

    if(!oEvidenceCollection.validateEvidenceCollection(m_pPolicy)) {
        fprintf(g_logFile, "permitAccess: Can't validate Evidence list\n");
        return false;
    }

    if(oEvidenceCollection.m_iNumEvidenceLists<1 || 
            oEvidenceCollection.m_rgiCollectionTypes[0]!=SIGNEDGRANT) {
        fprintf(g_logFile, "permitAccess: No Signed grant\n");
        return false;
    }
    pAssert= (SignedAssertion*) oEvidenceCollection.m_rgCollectionList[0]->m_rgEvidence[0];

    // map request to required access
    if(strcmp(req.m_szRequest, "createResource")==0)
        uVerb= MAYCREATE;
    else if(strcmp(req.m_szRequest, "sendResource")==0)
        uVerb= MAYWRITE;
    else if(strcmp(req.m_szRequest, "getResource")==0)
        uVerb= MAYREAD;
    else if(strcmp(req.m_szRequest, "getOwner")==0 || 
            strcmp(req.m_szRequest, "addOwner")==0 ||
            strcmp(req.m_szRequest, "removeOwner")==0)
        uVerb= MAYOWN;
    else if(strcmp(req.m_szRequest, "deleteResource")==0)
        uVerb= MAYDELETE;
    else {
        fprintf(g_logFile, "permitAccess: Unknown request\n");
        return false;
    }

#ifdef ACCESSTEST
    fprintf(g_logFile, "permitAccess: Checking assertions\n");
    fflush(g_logFile);
#endif

    assertionNode**  rgpAssertions= 
            (assertionNode**) malloc(sizeof(assertionNode*)*pAssert->m_iNumAssertions);
    pSaysPrincipal= m_pMeta->findPrincipal(pAssert->getPrincipalName());
#ifdef ACCESSTEST
    fprintf(g_logFile, "says principal name: %s\n", pAssert->getPrincipalName());
    fflush(g_logFile);
#endif
    if(pSaysPrincipal==NULL) {
        fprintf(g_logFile, "Who says?\n");
        return false;
    }
    for(i=0; i<pAssert->m_iNumAssertions; i++) {
        rgpAssertions[i]= new assertionNode();
        if(!rgpAssertions[i]->parseAssertion(pSaysPrincipal, pAssert->m_rgszAssertion[i], 
                                true)) {
            fprintf(g_logFile, "parseAssertion %d returned false\n", i);
            return false;
        }
    }

#ifdef ACCESSTEST
    fprintf(g_logFile, "permitAccess: Finished parsing %d assertions\n", pAssert->m_iNumAssertions);
    fflush(g_logFile);
#endif

    pSubjNode= m_Subjects.pFirst;
    while(pSubjNode!=NULL) {
        pSubjPrincipal= pSubjNode->pElement;
#ifdef ACCESSTEST
        fprintf(g_logFile, "Trying subject %s\n", pSubjPrincipal->m_szPrincipalName);
#endif
        for(i = 0; i < pAssert->m_iNumAssertions; i++) {
#ifdef ACCESSTEST
            fprintf(g_logFile, "trying assertion %d\n", i);
#endif
            if(rgpAssertions[i]->assertionSucceeds(pSubjPrincipal, uVerb, pResource,
                      pAssert->m_iNumAssertions, rgpAssertions)) {
#ifdef ACCESSTEST
                fprintf(g_logFile, "permitAccess: The assertion succeeds\n");
                fflush(g_logFile);
#endif
                return true;
            } 
            else {
#ifdef ACCESSTEST
                fprintf(g_logFile, "permitAccess: The assertion fails\n");
#endif
            }
        }
        pSubjNode= pSubjNode->pNext;
    }

    fprintf(g_logFile, "permitAccess: Finished checking assertion without finding one that succeeds\n");
    fflush(g_logFile);
    return false;
}


// ---------------------------------------------------------------------------


