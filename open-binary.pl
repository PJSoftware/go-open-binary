#!/bin/perl

use strict;

die "Must specify a file to open" unless $ARGV[0];

my $bin = $ARGV[0];
die "'$bin' not found" unless -f $bin;

print "Opening '$bin' in default software\n";
system qq{powershell -Command start '"$bin"'};
