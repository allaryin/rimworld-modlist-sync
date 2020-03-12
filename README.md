# RimWorld ModList Sync

This is a very rudimentary command line tool for synchronizing your current
[RimWorld](https://rimworldgame.com/) mod config with that of a save file. Eventually, it
might be extended to do other things - but there are no plans in the works.

Usage instructions and compiled binaries are pending my getting the tool into a minimum
viable state. But I do know that **it is probably not a good idea to edit the ModConfig.xml
while RimWorld is currently running** ;)

## Roadmap
* [x] Find RimWorld data directory
* [x] Read ModConfig.xml
* [ ] ... and Savegame.rws files
* [ ] Browse save files, possibly allowing fuzzy name matching
* [ ] Validate save vs config versions
* [ ] Copy modlist from a save into the config, saving a backup first
* [ ] Restore ModConfig.xml from a backup
