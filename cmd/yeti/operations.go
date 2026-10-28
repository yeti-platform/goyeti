package main

// What each operation does, worded for an agent deciding which to call. The
// CLI help and the MCP tool descriptions both use these, so they can't drift.
const (
	lookupDescription = "Look up values in Yeti: IPs, hostnames, URLs, hashes and other observables, " +
		"defanged or not. For each value, says whether Yeti knows it, its tags and context sources, " +
		"the entities (malware, threat actors, campaigns…) and observables linked to it, and the " +
		"indicators that match it. Start here with any indicator of compromise. Unknown values are " +
		"reported, never added."

	searchDescription = "Search Yeti by keyword across observables, entities, indicators and DFIQ, and by " +
		"meaning across entities, indicators and DFIQ. Use it for names and topics (\"Emotet\", " +
		"\"credential dumping\"); use lookup for observable values. Hits found by meaning come first, " +
		"with a 0-1 score."

	getDescription = "Get one Yeti object in full: every field, including context and tag expiry. Takes a " +
		"reference such as observables/123 or entities/45, as returned by lookup, search and neighbors."

	neighborsDescription = "List the objects linked to one Yeti object, through relationships such as " +
		"\"uses\" or \"c2\", up to 3 hops away. Use it to pivot: from a malware to its infrastructure, " +
		"from an IP to the actor behind it. Narrow with link types, target types and direction."

	yaraBundleDescription = "Build a compilable YARA source from YARA rules in Yeti, given by reference " +
		"(indicators/123) or by tag, with every rule they depend on."

	typesDescription = "List the object types this Yeti server knows, by family: observables, entities, " +
		"indicators and DFIQ. Use the names as target types for neighbors."
)

// The mutate tier's operations.
const (
	addObservablesDescription = "Add observables to Yeti: IPs, hostnames, URLs, hashes…, defanged or not. " +
		"Values Yeti already knows are left as they are, and every value gets the tags. Yeti guesses " +
		"each value's type unless one is given, and rejects values it doesn't recognize. Each value " +
		"reports created, exists or failed."

	addEntityDescription = "Add an entity to Yeti — a malware, threat actor, campaign, attack pattern, tool, " +
		"vulnerability… — unless one of that type and name exists. An existing entity's fields are " +
		"left as they are; the tags are added either way. \"yeti read types\" lists the entity types."

	addIndicatorDescription = "Add an indicator to Yeti — a YARA rule, Sigma rule, regex, query… — unless one " +
		"of that type and name exists. It needs a pattern and a diamond model vertex (adversary, " +
		"capability, infrastructure or victim); regex and query need a location. A YARA rule takes the " +
		"name of the rule in its pattern."

	linkDescription = "Link two Yeti objects with a relationship, e.g. an IP to the malware it serves as c2 " +
		"for, or an actor to the malware it uses. Takes the source's reference, the relationship type " +
		"and the target's reference. Objects already linked that way are left as they are."

	tagDescription = "Add tags to Yeti objects (observables, entities, indicators), given by reference. " +
		"Tag names are normalized: lowercase, spaces become underscores. Objects that have the tags " +
		"already are left unchanged."

	untagDescription = "Remove tags from Yeti objects, given by reference. The other tags are kept."

	addContextDescription = "Add context to a Yeti object: facts about it as name/value pairs, e.g. a note, " +
		"a first-seen date or a report URL. The context's source is \"agent:<name>\", the agent's " +
		"name, so Yeti's timeline shows which agent added it. Identical context isn't added twice."
)

// The admin tier's operations, and the read tier's view of tasks.
const (
	tasksDescription = "List Yeti's tasks — feeds, analytics, exports… — with whether each is enabled, " +
		"its status, last run and schedule. Narrow by name (any part of it, ignoring case) and by type."

	deleteDescription = "Delete Yeti objects, given by reference, with all their links. This is permanent: " +
		"do a dry run first, which says what each object is and how many links would go, then " +
		"confirm. Objects that don't exist are reported absent. Deleting a DFIQ object detaches its " +
		"children rather than deleting them."

	unlinkDescription = "Delete the link of a type from one Yeti object to another, leaving both objects. " +
		"Takes the source's reference, the relationship type and the target's reference, like link. " +
		"Needs confirmation; a link that doesn't exist is reported absent."

	runTaskDescription = "Queue a run of a Yeti task, such as a feed or an analytics, by its exact name. " +
		"The task must be enabled. Yeti runs it in the background: list the task with tasks to follow " +
		"its status and last run. Oneshot analytics act on one observable Yeti knows: pass its value as " +
		"the \"value\" parameter."

	enableTaskDescription = "Enable Yeti tasks by name, so that Yeti runs them on schedule and on demand. " +
		"Tasks that are enabled already are left unchanged."

	disableTaskDescription = "Disable Yeti tasks by name, so that Yeti stops running them. Tasks that are " +
		"disabled already are left unchanged."
)
