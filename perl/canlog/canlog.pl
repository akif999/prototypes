package main;

use 5.008001;
use strict;
use warnings;
use utf8;

my $log = Log->new;
$log->parse("   0.002763 1  1E5             Rx   d 8 4C 00 21 10 00 00 00 B9  Length = 228000 BitCount = 118 ID = 485");
print $log->string;

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
    $self->{TIME}   = undef;
    $self->{CH}     = undef;
    $self->{ID}     = undef;
    $self->{DIR}    = undef;
    $self->{STAT}   = undef;
    $self->{DLC}    = undef;
    $self->{DATA}   = [];
    $self->{REMAIN} = [];
    return bless $self, $class;
}

sub parse {
    my ($self, $line) = @_;
    my @fields = split(/\s+/, $line);
    my ($time, $ch, $id, $dir, $stat, $dlc) = @fields[1..7];
    my @data = @fields[7..6+$dlc];
    my @remain = @fields[6+$dlc..scalar(@fields)-1];

    $self->{TIME}   = int($time * 1000000);
    $self->{CH}     = int($ch);
    $self->{ID}     = hex($id);
    $self->{DIR}    = $dir eq "Rx" ? RX : TX;
    $self->{STAT}   = $stat;
    $self->{DLC}    = $dlc;
    @{$self->{DATA}}   =  map(hex, @data);
    $self->{REMAIN} = @remain;
}

sub string {
    my $self = shift;
    my $str .= sprintf "%f %d, %03X %s %X", $self->{TIME} / 1000000, $self->{CH}, $self->{ID}, $self->{DIR} eq RX ? "Rx" : "Tx", $self->{DLC};
    $str .= sprintf " %02X", $_ foreach @{$self->{DATA}};
    return $str
}

1;
