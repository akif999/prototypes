use strict;
use warnings;
use Data::Dumper;
use Spreadsheet::ParseXLSX;

my $parser = Spreadsheet::ParseXLSX->new;
my $workbook = $parser->parse("testcalc.xlsx");
if ( !defined $workbook ) {
    die $parser->error(), ".\n"
}

my $worksheet = $workbook->worksheet(0);
my ( $row_min, $row_max ) = $worksheet->row_range();
my ( $col_min, $col_max ) = $worksheet->col_range();

my $onep = new_person();
my @persons = ();

for my $row ( $row_min .. $row_max ) {
        my $Name ={};
        my $Child = ();

    for my $col ( $col_min .. $col_max ) {

        my $cell = $worksheet->get_cell( $row, $col );
        if ( $cell->value() ne "" ) {

            if ( $col == 0 ) {
                $onep->{Name} = $cell->value();
            } elsif ( $col == 1 ) {
                $onep->{Comment} = $cell->value();
            } elsif ( $col == 2 ) {
                $Name = {Name => $cell->value()};
            } elsif ( $col == 3 ) {
                my @list = split /,/, $cell->value();
                foreach my $parts (@list) {
                    push @$Child, {Name => $parts};
                }
            } else {
                # pass others
            }
        }
    }
    push @{$onep->{Monsters}}, {$Name->{Name}, $Child};

    my $nextcell = $worksheet->get_cell( $row + 1, 0 );
    if ( $row == $row_max ) {
        push @persons, $onep;
    } elsif ( $nextcell->value ne "") {
        push @persons, $onep;
        $onep = new_person();
    } else {

    }
}

print Dumper \@persons;

sub new_person {
    my $person = {
        Name     => '',
        Comment  => '',
        Monsters => [
    #        {
    #            Name  => '',
    #            Child => [
    #                Name => ''
    #            ],
    #        }
        ]
    };
    return $person;
}
