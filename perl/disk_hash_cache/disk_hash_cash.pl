use strict;
use warnings;
use utf8;
use Storable qw/nfreeze nstore thaw retrieve/;
use Spreadsheet::ParseXLSX;

my $file = $ARGV[0];
my $database = ();

if (-e 'cache.bin') {
    $database = retrieve 'cache.bin';
} else {
    parse_personal_database(\@$database, $file);
    nstore $database, 'cache.bin';
}
print join("\n",  @$database), "\n";

sub parse_personal_database {
    my ($db, $file) = @_;

    my $parser   = Spreadsheet::ParseXLSX->new();
    my $workbook = $parser->parse($file);

    if (!defined $workbook) {
        die $parser->error(), "\n";
    }

    for my $worksheet ($workbook->worksheets()) {
        my ($row_min, $row_max) = $worksheet->row_range();
        my ($col_min, $col_max) = $worksheet->col_range();

        for my $row ($row_min .. $row_max) {
            my $str = "";
            for my $col ($col_min .. $col_max) {
                my $cell = $worksheet->get_cell($row, $col);
                $str .= $cell->value." ";
            }
            push @$db, $str;
        }
    }
}
