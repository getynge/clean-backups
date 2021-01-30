# Overview
clean-backups is a rudimentary tool for personal use.

clean-backups takes exactly one argument: the folder to clean.

clean-backups will delete any file for which all of the following are true:
* the files name, sans the format, is exactly 12 digits (a timestamp YYYYMMDDhhmm)
* the file format is .tar.gz
* the name represents a timestamp where the hour is _not_ 23
* the name represents a timestamp that is more than 24 hours in the past

clean-backups will also delete files where the following is true:
* the name represents a timestamp that is more than 30 days in the past

clean backups will search the entire directory tree beneath the supplied folder.

This utility should only be used on file trees consisting exclusively of backups that
are intended to be no more than a month old.