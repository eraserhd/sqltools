{ nixpkgs ? (import ./nixpkgs.nix), ... }:
let
  pkgs = import nixpkgs {
    config = {};
    overlays = [
      (import ./overlay.nix)
    ];
  };
in {
  test = pkgs.runCommandNoCC "sqltools-test" {} ''
    mkdir -p $out
    : ${pkgs.sqltools}
  '';
}