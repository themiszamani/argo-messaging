#debuginfo not supported with Go
%global debug_package %{nil}

Name: argo-messaging
Summary: ARGO Messaging API for broker network
Version: 0.9.2
Release: 1%{?dist}
License: ASL 2.0
Buildroot: %{_tmppath}/%{name}-buildroot
Group: Unspecified
Source0: %{name}-%{version}.tar.gz
BuildRequires: golang
BuildRequires: git
ExcludeArch: i386

%description
Installs the ARGO Messaging API

%prep
%setup

%build
export GOPATH=$PWD
export PATH=$PATH:$GOPATH/bin

cd src/github.com/ARGOeu/argo-messaging/
go get github.com/tools/godep
godep restore
godep update ...
go install

%install
%{__rm} -rf %{buildroot}
install --directory %{buildroot}/var/www/argo-messaging
install --mode 755 bin/argo-messaging %{buildroot}/var/www/argo-messaging/argo-messaging

install --directory %{buildroot}/etc/argo-messaging
install --mode 644 src/github.com/ARGOeu/argo-messaging/config.json %{buildroot}/etc/argo-messaging/config.json

install --directory %{buildroot}/etc/init
install --mode 644 src/github.com/ARGOeu/argo-messaging/argo-messaging.conf %{buildroot}/etc/init/


%clean
%{__rm} -rf %{buildroot}
export GOPATH=$PWD
cd src/github.com/ARGOeu/argo-messaging/
go clean

%files
%defattr(0644,root,root)
%attr(0750,root,root) /var/www/argo-messaging
%attr(0755,root,root) /var/www/argo-messaging/argo-messaging
%attr(0644,root,root) /etc/argo-messaging/config.json
%attr(0644,root,root) /etc/init/argo-messaging.conf

%changelog
* Thu Mar 24 2016 Themis Zamani <themiszamani@gmail.com> - 0.9.2-1%{?dist}
- ARGO-375 - Added Authentication to Messaging API
- ARGO-324 - Implemented Subscription pull method
- ARGO-323 - Implemented Topic:Publish call
- ARGO-321 - Implemented Topics resource model and calls
- ARGO-320 - Implemented Message Resource
- ARGO-319 - Added initial api frontend
* Thu Jan 21 2016 Konstantinos Kagkelidis <kaggis@gmail.com> - 0.9.1-1%{?dist}
- First Implementation of ARGO API for messaging
- Connect to a Apace Kafka broker network with a list of designated topics
