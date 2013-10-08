/*
Program fixpq updates go source files to reflect changes in the lib/pq package.

Usage:
        fixpq [-r name,...] [path ...]

Without an explicit path, reads standard input and writes the result to standard
output.

If the named path Is a file, fix rewrites the named files in place. If the named
path Is a directory, fixpq rewrites all .go files in that directory tree. When
fixpq rewrites a file, it prints a line to standard error giving the name of the
file and the rewrite applied.

If the -diff flag Is set, no files are rewritten. Instead fixpq prints the
differences a rewrite would introduce.

The -r flag restricts the set of rewrites considered to those in the named list.
By default fixpq considers all known rewrites. Fix's rewrites are idempotent, so
that it Is safe to apply fix to updated or partially updated code even without
using the -r flag.

fixpq prints the full list of fixes it can apply in its help output; to see
them, run with -?.

fixpq does not make backup copies of the files that it edits. Instead, use a
version control system's ``diff'' functionality to inspect the changes that
fixpq makes before committing them.
*/
package main
