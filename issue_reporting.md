# Issue Reporting

This guides you through submitting a bug report for CloudQuery. Following these guidelines helps maintainers and the community understand your report :pencil:, reproduce the behavior :computer: :cloud:, and find related reports :mag_right:.	


## Reporting Bugs

- [ ] **Determine the right repository**: If it's a general issue or an issue with an officially-supported plugin, please open the bug in [cloudquery/cloudquery](https://github.com/cloudquery/cloudquery). If it's an issue with a community plugin, please open the bug in the relevant plugin repository. 	
- [ ] **Search for possible duplicates**: If you see an already open issue, feel free to comment and add additional information. If the issue is closed please open a new one and link to the closed.	
- [ ] **Fill in the bug report template**: Try to fill in as much as possible with many details using the [bug report template](https://github.com/cloudquery/cloudquery/tree/main/.github/ISSUE_TEMPLATE/bug_report.md)	

### How Do I Submit a (Good) Bug Report? 
Explain the problem and include additional details to help maintainers reproduce the problem:

* **Use a clear and descriptive title** for the issue to identify the problem
* **Describe the Bug** in as many details as possible. For example, start by explaining how and where you are running CloudQuery (local machine, cloud service, docker, k8s, CI Pipeline, etc)
* **Provide specific examples to demonstrate the steps**. Include links to GitHub gists and or files, or copy/pasteable snippets to help give context to the issue. If you're providing snippets in the issue, use [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines)
* **Explain which behavior you expected to see instead and why.**
* **Include (sanitized) log output** execute CloudQuery with the `--enable-console-log` and `-v` flags to get all of the debug information

If possible, include details about your configuration and environment:

* **Which version of CloudQuery are you using?** You can get the exact version by running `cloudquery version`
* **What's in your cloudquery.yml**? Include as much of the `cloudquery.yml` as possible. This will allow the community to work to reproduce the issue and identify workarounds and/or create fixes to your issues


## Suggesting Enhancements

- [ ] **Determine the right repository**: If it's a general enhancement, or an enhancement to an official plugin, please open the issue in [cloudquery/cloudquery](https://github.com/cloudquery/cloudquery).	
- [ ] **Search for possible duplicates**: If you see an already open issue, feel free to comment and add additional information. If the issue is closed please open a new one and link to the closed.	
- [ ] **Fill in the feature request template**: Try to fill-in as much as possible with many details the [feature request template](https://github.com/cloudquery/cloudquery/tree/main/.github/ISSUE_TEMPLATE/feature_request.md)	

### How Do I Submit a (Good) Enhancement Suggestion?

* **Use a clear and descriptive title** for the issue to identify the suggestion
* **Describe the problem** In detail please try and convey the workflow or functionality you are trying to implement. This will help the community design and implement tooling that is both intuitive to use across many different domains as well as applicable to you and your specific challenge
* **Describe the use case** for this feature in as much detail as possible. Be sure to include any relevant information including links or other implementations


## Requesting Resources or New Plugins

- [ ] **Determine the right repository**: If it is a new plugin, or fits with an officially-supported plugin, please open the issue in [cloudquery/cloudquery](https://github.com/cloudquery/cloudquery). If it's a resource request for a community plugin, please open the issue against the relevant plugin repository. 	
- [ ] **Search for possible duplicates**: If you see an already open issue, feel free to comment and add additional information. If the issue is closed please open a new one and link to the closed.	
- [ ] **Fill in the new resource template**: Try to fill in as much as possible with many details using the [new resource template](https://github.com/cloudquery/cloudquery/tree/main/.github/ISSUE_TEMPLATE/new_resource.md)	
