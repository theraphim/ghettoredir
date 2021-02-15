Name:		ghettoredir
Version:	0.0.1
Release:	1%{?dist}
Summary:	gitbucket go get helper
License:	MIT
URL:		http://stingr.net
Source0:	%{name}-%{version}.tar.gz

BuildRequires:	golang
BuildRequires:  systemd
%{?systemd_requires}

%post
%systemd_post %{name}.service %{name}.socket

%preun
%systemd_preun %{name}.service %{name}.socket

%postun
%systemd_postun_with_restart %{name}.service %{name}.socket

%description
gitbucket go get helper

%prep
%setup -q

%build
mkdir -p goapp/src/github.com/theraphim goapp/bin
ln -s ${PWD} goapp/src/github.com/theraphim/%{name}
export GO111MODULE=off
export GOPATH=${PWD}/goapp
%gobuild -o goapp/bin/%{name} github.com/theraphim/%{name}

%install

%{__install} -d $RPM_BUILD_ROOT%{_bindir}
%{__install} -v -D -t $RPM_BUILD_ROOT%{_bindir} goapp/bin/%{name}
%{__install} -d $RPM_BUILD_ROOT%{_unitdir}
%{__install} -v -D -t $RPM_BUILD_ROOT%{_unitdir} %{name}.service
%{__install} -v -D -t $RPM_BUILD_ROOT%{_unitdir} %{name}.socket
%{__install} -d -m 0755 %{buildroot}%{_sysconfdir}/%{name}
%{__install} -d $RPM_BUILD_ROOT%{_sysconfdir}/sysconfig
%{__install} -m 644 -T %{name}.sysconfig %{buildroot}%{_sysconfdir}/sysconfig/%{name}
%{__install} -d -m 0755 %{buildroot}%{_sharedstatedir}/%{name}

%files
%{_bindir}/%{name}
%dir %attr(-,%{name},%{name}) %{_sharedstatedir}/%{name}
%{_unitdir}/%{name}.service
%{_unitdir}/%{name}.socket
%config(noreplace) %{_sysconfdir}/%{name}
%config(noreplace) %{_sysconfdir}/sysconfig/%{name}


%changelog
