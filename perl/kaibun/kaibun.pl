use strict;
use warnings;
use utf8;

my @answers = kaibunparse($ARGV[0]);
foreach my $answer (@answers) {
    printf "answer: bin = %b, oct = %o, dec = %d\n", $answer, $answer, $answer;
}


sub kaibunparse {
    my $max = shift;

    my @answers = ();
    for (my $i=1; $i<=$max ;$i++) {
        my $bin = sprintf("%b", $i);
        my $oct = sprintf("%o", $i);
        my $dec = sprintf("%d", $i);

        my $iskaibun = splitandcompare($bin, $oct, $dec);
        if ($iskaibun eq "true") {
            push @answers, $i;
        }
    }
    return @answers;
}

sub splitandcompare {
    my @numbers = @_;

    foreach my $num (@numbers)  {
        if (length($num) == 1 ) {
            next;
        }

        if (length($num) % 2 != 0) {
            substr($num, (length($num) / 2), 1) = "";
        }

        my $len = length($num) / 2;
        (my $pre, my $post) = $num =~ /.{$len}/g;
        if ($pre eq reverse($post)) {
            next;
        } else {
            return "false";
        }
    }
    return "true";
}
