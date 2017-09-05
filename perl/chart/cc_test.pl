use strict;
use warnings;
use utf8;

use Chart::Clicker;
use Chart::Clicker::Data::Series;
use Chart::Clicker::Data::DataSet;

my $cc1 = Chart::Clicker->new;
my $cc2 = Chart::Clicker->new;

my $series1 = Chart::Clicker::Data::Series->new(
    keys   => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    values => [0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0],
);


my $series2 = Chart::Clicker::Data::Series->new(
    keys   => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    values => [0, 500, 310, 720, 500, 180, 120, 90, 320, 700, 880],
);

my $ds1 = Chart::Clicker::Data::DataSet->new(series => [ $series1 ]);
my $ds2 = Chart::Clicker::Data::DataSet->new(series => [ $series2 ]);

$cc1->add_to_datasets($ds1);
$cc1->draw;
$cc1->write('hoge.png');

$cc2->add_to_datasets($ds2);
$cc2->draw;
$cc2->write('fuga.png');
