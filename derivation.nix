{ buildGoModule, lib, fetchFromGitHub, ... }:

buildGoModule rec {
  pname = "sqltools";
  version = "0.1.0";

  src = fetchFromGitHub {
    owner = "eraserhd";
    repo = pname;
    rev = "v${version}";
    sha256 = "";
  };

  vendorSha256 = "pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";

  meta = with lib; {
    description = "Tools for parsing and refactoring SQL";
    homepage = "https://github.com/eraserhd/sqltools";
    license = licenses.publicDomain;
    platforms = platforms.all;
    maintainers = [ maintainers.eraserhd ];
  };
}
