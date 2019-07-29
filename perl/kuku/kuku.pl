foreach my $x (@{[1..9]}) {
    foreach my $y (@{[1..9]}) {
        printf "%02d ", $x * $y;
    }
    print "\n";
}
