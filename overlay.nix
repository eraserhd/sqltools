self: super: {
  sqltools = super.callPackage ./derivation.nix {
    fetchFromGitHub = _: ./.;
  };
}
