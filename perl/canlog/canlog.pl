package main;

use 5.008001;
use strict;
use warnings;
use utf8;

use Path::Tiny;
use Getopt::Kingpin;

my $kingpin = Getopt::Kingpin->new;

my $filename = $kingpin->arg("file", "input log file")->string;

$kingpin->parse;

my $path = path($filename);

foreach my $line ($path->lines) {
    chomp($line);
    if ($line =~ /^\s+\d+\.\d+\s[1-9]/) {
        my $log = Log->new;
        $log->parse($line);
        print $log->string, "\n";
    }
}

package Log;

use 5.008001;
use strict;
use warnings;

use constant {
    RX => 0,
    TX => 1,
};

sub new {
    my $class = shift;
    my $self = {};
    $self->{TIME}    = undef;
    $self->{CH}      = undef;
    $self->{ID}      = undef;
    $self->{DIR}     = undef;
    $self->{STAT}    = undef;
    $self->{DLC}     = undef;
    $self->{DATA}    = [];
    $self->{REMAIN}  = [];
    $self->{SPLITER} = " ";
    return bless $self, $class;
}

sub parse {
    my ($self, $line) = @_;
    my @fields = split(/\s+/, $line);
    my ($time, $ch, $id, $dir, $stat, $dlc) = @fields[1..7];
    my @data = @fields[7..6+$dlc];
    my @remain = @fields[6+$dlc..scalar(@fields)-1];

    $self->{TIME}    = int($time * 1000000);
    $self->{CH}      = int($ch);
    $self->{ID}      = hex($id);
    $self->{DIR}     = $dir eq "Rx" ? RX : TX;
    $self->{STAT}    = $stat;
    $self->{DLC}     = $dlc;
    @{$self->{DATA}} =  map(hex, @data);
    $self->{REMAIN}  = @remain;
}

sub string {
    my $self = shift;
    my $sp = $self->{SPLITER};
    my $str .= sprintf "%f%s%d%s%03X%s%s%s%X", $self->{TIME} / 1000000, $sp, $self->{CH},
        $sp, $self->{ID}, $sp, $self->{DIR} eq RX ? "Rx" : "Tx", $sp, $self->{DLC};
    $str .= sprintf "%s%02X", $sp, $_ foreach @{$self->{DATA}};
    return $str
}

sub spliter {
    my ($self, $sp) = @_;
    $self->{SPLITER} = $sp;
}

1;
