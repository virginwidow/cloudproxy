%{!?luaver: %global luaver %(lua -e "print(string.sub(_VERSION, 5))")}
%global	luapkgdir %{_datadir}/lua/%{luaver}

Name:		datalog
Version:	2.4
Release:	1%{?dist}

Source0:	http://downloads.sourceforge.net/datalog/%{name}-%{version}.tar.gz

URL:		http://sourceforge.net/projects/datalog
BuildRoot:	%{_tmppath}/%{name}-%{version}-%{release}-root

Summary:	A Lightweight Deductive Database using Datalog
License:	LGPLv2+
%if 0%{?rhel}
Group:		Applications/Databases
%endif

BuildRequires:	texinfo, lua-devel < 5.3

%description
This package contains a lightweight deductive database system.
Queries and database updates are expressed using Datalog--a
declarative logic language in which each formula is a function-free
Horn clause, and every variable in the head of a clause must appear in
the body of the clause.  The use of Datalog syntax and an
implementation based on tabling intermediate results, ensures that all
queries terminate.

The components in this package are designed to be small, and usable on
memory constrained devices.  The package includes an interactive
interpreter for Datalog, and the development package has a library
that can be used to embed the interpreter into C programs.

%package devel
Summary: Datalog header file and library
%if 0%{?rhel}
Group: Development/Libraries
%endif
Requires: datalog = %{version}

%description devel
This package includes the header file and library that can be used to
embed a datalog interpreter into C programs.

%prep
%setup -q

%build
%configure --with-lua --enable-shared --disable-static
make %{?_smp_mflags}

%install
make DESTDIR=%{buildroot} install
mkdir -p %{buildroot}%{luapkgdir}
install -m 644 %{name}.lua %{buildroot}%{luapkgdir}
rm -rf %{buildroot}/%{_libdir}/lib%{name}.la
rm -rf %{buildroot}/%{_datadir}/%{name}
rm -rf %{buildroot}/%{_infodir}/dir

%clean
rm -rf %{buildroot}

%post
/sbin/ldconfig
if [ -f %{_infodir}/%{name}.info.gz ]; then
  /sbin/install-info %{_infodir}/%{name}.info.gz %{_infodir}/dir || :
fi

%postun
/sbin/ldconfig
if [ $1 -eq 0 ]; then
  if [ -f %{_infodir}/%{name}.info.gz ]; then
    /sbin/install-info --delete %{_infodir}/%{name}.info.gz %{_infodir}/dir || :
  fi
fi

%files
%defattr (-, root, root, -)
%doc %{name}.html ChangeLog README COPYING.LIB AUTHORS NEWS
%{_bindir}/%{name}
%{_libdir}/lib%{name}.so.*
%{_infodir}/%{name}.info.gz
%{_mandir}/man1/*
%{luapkgdir}/%{name}.lua

%files devel
%defattr (-, root, root, -)
%{_includedir}/%{name}.h
%{_libdir}/lib%{name}.so

%changelog
* Wed Feb 13 2013 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 2.3-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_19_Mass_Rebuild

* Mon Oct 22 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.3-2
- Remove lua(api) requirement

* Thu Jun  7 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.2-4
- Added lua(abi) requirement

* Tue Jun  5 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.2-3
- Removed rm of build root in %%install
- Group field defined only for RHEL
- Removed license field in subpackage devel

* Fri May 25 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.2-2
- Changed %%define to %%global

* Thu Apr 26 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.2-1
- Changed devel requires from libdatalog to datalog

* Sun Apr  8 2012 John D. Ramsdell <ramsdell@mitre.org> - 2.1-1
- Added AUTHORS and NEWS to %%doc

* Thu Jan 26 2012 John D. Ramsdell <ramsdell@mitre.org> - 1.8-2
- Use lua to determine its version number

* Wed Jan 18 2012 John D. Ramsdell <ramsdell@mitre.org> - 1.8-1
- Added a manual page

* Fri Jan 13 2012 John D. Ramsdell <ramsdell@mitre.org> - 1.7-3
- Fix devel license and summary
- chmod 644 on installed datalog.lua

* Wed Oct 12 2011 John D. Ramsdell <ramsdell@mitre.org> - 1.7-2
- Added devel package
- Moved luapkgdir def to top of file
- sf.net --> sourceforge.net
- Removed Packager field
- Removed newline between %%description and text

* Thu Sep 29 2011 John D. Ramsdell <ramsdell@mitre.org> - 1.7-1
- Use installed Lua package and shared libraries

* Tue Sep 20 2011 John D. Ramsdell <ramsdell@mitre.org> - 1.6-2
- Removed devel package

* Sat Aug 6 2011 John D. Ramsdell <ramsdell@mitre.org> - 1.6-1
- Fixed license name by adding version number
- Dropped vendor field
- Switched to standard name for build root
- %%defaddrs now contain four fields
- Removed asterisks in %%files
- Added COPYING.LIB

* Tue Jul 7 2011 John D. Ramsdell <ramsdell@mitre.org> - 1.4-1
- Initial spec release
