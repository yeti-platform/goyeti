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
