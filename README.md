
This is a simple launcher to start commands by whistling. It uses sndpeek (http://soundlab.cs.princeton.edu/software/sndpeek/)
and is inspired (a lot) by http://www.ibm.com/developerworks/library/os-whistle/index.html?ca=dgr-lnxw97whistlework

== Structure of the repo ==

	* sndlib contains the core of the program, including the note detector and sequence detector.
		most constants are still in the files
	* persistance contains primitives to save the configuration to file
	* calibrate.go is the program that inputs your commands and whistle
	* whistlelauncher parse your config file and launch the program

== TODO ==

* Actually launch the program (we're just matching the right command now)
* Document the code
* Extract constants from the library files