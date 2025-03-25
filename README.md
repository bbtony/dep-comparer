# dep-comparer

is a small utility for developers and others who need to create a list of dependencies from various package manager files. 

Currently, it supports Golang and its go.mod files. I plan to add support for other languages in the near future.

### How it works?

The main idea is to provide a "pivot" table of your dependencies across all your repositories. This is particularly useful when you're starting work on the Software Development Lifecycle (SDLC) and need to analyze dependencies in the initial stages. 

For example, if you want to retrieve all packages and their versions from your Golang go.mod files, 
you can use the following command:
```bash
dep-comparer testdata/go1.mod testdata/go2.mod testdata/go3.mod
```
The output will be a report in CSV format, for example: [examples/report-1736365627.csv](examples/report-1736365627.csv)

### Experimental feature

If you examine the report generated from the testdata go.mod files, you'll notice a large number of dependencies. 
You might consider visualization, and you'd be right. 
dep-comparer supports an experimental feature: generating a report in [dot-format](https://en.wikipedia.org/wiki/DOT_(graph_description_language)) for [graphviz](https://graphviz.org/Gallery/directed/) with Graphviz. 
To generate a DOT report, use the `-dot` flag, like this:
```bash
dep-comparer -dot testdata/go1.mod testdata/go2.mod testdata/go3.mod
```

This feature requires the Graphviz dependency, which you'll need to install in your environment. Once installed, you can use the DOT report. 
For example:
```bash
sfdp -Gsize=67! -Goverlap=prism -Tsvg examples/graph_1736365627.dot > graph_1736365627.dot/root.svg
```

**IMPORTANT:** 
I'm not a Graphviz or visualization expert, so feel free to use, modify, and extend the DOT report as needed.
:-)

![](examples/root.svg)