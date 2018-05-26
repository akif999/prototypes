use strict;
use warnings;
use YAML::Tiny;
use Data::Dumper;

my $data = YAML::Tiny->new;
$data = YAML::Tiny->read('test.yaml');
print Dumper($data);
