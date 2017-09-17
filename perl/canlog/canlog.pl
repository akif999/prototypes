package main;

use 5.008001;
use strict;
use warnings;
use utf8;

use Getopt::Kingpin;

my $kingpin = Getopt::Kingpin->new;
my $filename = $kingpin->arg("file", "input log file")->string;
$kingpin->parse;

my $canlog = CanLog->new;
$canlog->parse($filename);
print $canlog->string . "\n";

$canlog->stringer(sub{
    my ($obj, $sp) = @_;
    my $str .= sprintf "%f%s%03X", $obj->{TIME} / 1000000, $sp, $obj->{ID};
    $str .= sprintf "%s%02X", $sp, $_ foreach @{$obj->{DATA}};
    return $str
});
print $canlog->string . "\n";

my $timediff = $canlog->timediff;
my $strings  = $canlog->strings;
for (my $i = 0; $i< scalar(@$strings); $i++) {
    printf "%f %s\n", @$timediff[$i] / 1000000, @$strings[$i];
}

package CanLog;

use 5.008001;
use strict;
use warnings;
use Path::Tiny;

sub new {
    my $class = shift;
    my $self = {};
    $self->{LOGS} = ();
    $self->{SPLITER}  = " ";
    $self->{STRINGER} = \&string_simple;
    return bless $self, $class;
}

sub parse {
    my ($self, $filename) = @_;
    my $path = path($filename);
    foreach my $line ($path->lines) {
        chomp($line);
        if ($line =~ /^\s+\d+\.\d+\s[1-9]/) {
            my $log = Log->new;
            $log->parse($line);
            push @{$self->{LOGS}}, $log;
        }
    }
}

sub timediff {
    my $self = shift;
    my @diff = ();
    my $prev = 0000000;
    foreach my $log (values($self->{LOGS})) {
        push @diff, $log->{TIME} - $prev;
        $prev = $log->{TIME};
    }
    return \@diff;
}

sub strings {
    my $self = shift;
    my @strings = ();
    foreach my $log (values($self->{LOGS})) {
        push @strings, $log->string($self->{STRINGER}, $self->{SPLITER});
    }
    return \@strings;
}

sub string {
    my $self = shift;
    my $string = "";
    foreach my $log (values($self->{LOGS})) {
        $string .= $log->string($self->{STRINGER}, $self->{SPLITER}) . "\n";
    }
    return $string;
}

sub string_simple {
    my ($obj, $sp) = @_;
    my $str .= sprintf "%f%s%d%s%03X%s%s%s%X", $obj->{TIME} / 1000000, $sp, $obj->{CH},
        $sp, $obj->{ID}, $sp, $obj->{DIR} eq Log->RX ? "Rx" : "Tx", $sp, $obj->{DLC};
    $str .= sprintf "%s%02X", $sp, $_ foreach @{$obj->{DATA}};
    return $str
}

sub spliter {
    my ($self, $sp) = @_;
    $self->{SPLITER} = $sp;
}

sub stringer {
    my ($self, $subroutine) = @_;
    $self->{STRINGER} = $subroutine;
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
    $self->{TIME}     = undef;
    $self->{CH}       = undef;
    $self->{ID}       = undef;
    $self->{DIR}      = undef;
    $self->{STAT}     = undef;
    $self->{DLC}      = undef;
    $self->{DATA}     = [];
    $self->{REMAIN}   = [];
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
    my ($self, $stringer, $spliter) = @_;
    return &{$stringer}($self, $spliter);
}
