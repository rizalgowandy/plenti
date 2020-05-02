package generated

// Do not edit, this file is automatically generated.

// Defaults: scaffolding used in 'new site' command
var Defaults = map[string][]byte{
	"/.gitignore": []byte(`public
node_modules`),
	"/content/blog/_blueprint.json": []byte(`{
    "title": "text",
    "body": ["text"],
    "author": "text",
    "date": "date"
}`),
	"/content/blog/adding_pletiform.json": []byte(`{
    "title": "Build sites with good form",
    "body": [
        "Need an easy webform solution?",
        "Try adding a <a href='https://plentiform.com' target='blank' rel='noopener noreferrer'>plentiform</a>! (Coming soon)"
    ],
    "author": "Jim Fisk",
    "date": "1/26/2020"
}`),
	"/content/blog/post1.json": []byte(`{
    "title": "Post 1",
    "body": [
        "The first of the posts"
    ],
    "author": "Jim Fisk",
    "date": "1/24/2020"
}`),
	"/content/blog/post2.json": []byte(`{
    "title": "Post 2",
    "body": [
        "This is the second post."
    ],
    "author": "Jim Fisk",
    "date": "1/25/2020"
}`),
	"/content/index.json": []byte(`{
	"title": "My Plenti Site",
	"intro": {
		"slogan": "Visit the <a href=\"https://svelte.dev/tutorial\">Svelte tutorial</a> to learn how to build Svelte apps.",
		"color": "red"
	}
}`),
	"/content/pages/_blueprint.json": []byte(`{
	"title": "text",
	"description": ["text"],
	"author": "text"
}`),
	"/content/pages/about.json": []byte(`{
	"title": "About Plenti",
	"description": [
		"Plenti is <a href='https://jamstack.org/' target='blank' rel='noopener noreferrer'>JAMstack</a> framework with a modern frontend for creating dynamic experiences. We've cut out as many dependencies as possible so you can focus on being productive instead of wrestling with a complicated toolchain.",
		"The <a href='https://svelte.dev/' target='blank' rel='noopener noreferrer'>Svelte</a> frontend <em>cuts weight</em> so users get a snappy experience, even with bad internet connections or underpowered devices.",
		"The <a href='https://golang.org/' target='blank' rel='noopener noreferrer'>Go</a> backend <em>cuts wait</em> so apps build faster allowing devs to get more done and editors to get realtime feedback on content changes.",
		"Thanks for taking a look!"
	],
	"author": "Jim Fisk"
}`),
	"/content/pages/contact.json": []byte(`{
	"title": "Contact",
	"description": [
		"The project is 100% open source, so if you'd like to fork it for your own purposes, or help us out by reporting bugs / contributing code: <a href=\"https://github.com/plentico/plenti\">https://github.com/plentico/plenti</a>"
	],
	"author": "Jim Fisk"
}`),
	"/layout/components/grid.svelte": []byte(`<script>
  import { sortByDate } from '../scripts/sort_by_date.svelte';
  export let items, filter;
</script>

<div class="grid">
  {#each sortByDate(items) as item}
		{#if item.type == filter}
      <a class="grid-item" href="{item.path}">{item.fields.title}</a>
		{/if}
  {/each}
</div>

<style>
  .grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-column-gap: 10px;
    grid-row-gap: 10px;
  }
  .grid-item {
    display: flex;
    flex-grow: 1;
    height: 200px;
    background: var(--base);
    align-items: center;
    justify-content: center;
  }
</style>
`),
	"/layout/content/404.svelte": []byte(`<h1>Oops... 404 not found</h1>
<a href="/">Go home?</a>`),
	"/layout/content/blog.svelte": []byte(`<script>
	export let title, body, author, date;
</script>

<h1>{title}</h1>

<p><em>{#if author}Written by {author}{/if}{#if date}&nbsp;on {date}{/if}</em></p>

<div>
  {#each body as paragraph}
    <p>{@html paragraph}</p>
  {/each}
</div>

<details>
  <summary>Uses the "Blog" template</summary>
  <pre><code>layout/content/blog.svelte</code></pre>
</details>

<p><a href="/">Back home</a></p>
`),
	"/layout/content/index.svelte": []byte(`<script>
	export let title, intro, allNodes;
	import Grid from '../components/grid.svelte';
</script>

<h1>{title}</h1>

<section id="intro">
	<p>{@html intro.slogan}</p>
</section>

<div>
	<h3>Recent blog posts:</h3>
	<Grid items={allNodes} filter="blog" />
	<br />
</div>

<details>
  <summary>Uses the "Index" template</summary>
  <pre><code>layout/content/index.svelte</code></pre>
</details>
`),
	"/layout/content/pages.svelte": []byte(`<script>
	export let title, description;
</script>

<h1>{title}</h1>

<div>
  {#each description as paragraph}
    <p>{@html paragraph}</p>
  {/each}
</div>

<details>
  <summary>Uses the "Pages" template</summary>
  <pre><code>layout/content/pages.svelte</code></pre>
</details>

<p><a href="/">Back home</a></p>
`),
	"/layout/ejected/build.js": []byte(`import svelte from 'svelte/compiler.js';
import 'svelte/register.js';
import relative from 'require-relative';
import path from 'path';
import fs from 'fs';

// Get the arguments from Go command execution.
const args = process.argv.slice(2)

// -----------------
// Helper Functions:
// -----------------

// Create any missing sub folders.
const ensureDirExists = filePath => {
	let dirname = path.dirname(filePath);
	if (fs.existsSync(dirname)) {
		return true;
	}
	ensureDirExists(dirname);
	fs.mkdirSync(dirname);
}

// Concatenates HTML strings together.
const injectString = (order, content, element, html) => {
	if (order == 'prepend') {
		return html.replace(element, content + element);
	} else if (order == 'append') {
		return html.replace(element, element + content);
	}
};

// -----------------------
// Start client SPA build:
// -----------------------

let clientBuildStr = JSON.parse(args[0]);

clientBuildStr.forEach(arg => {

	let layoutPath = path.join(path.resolve(), arg.layoutPath)
	let component = fs.readFileSync(layoutPath, 'utf8');

	// Create component JS that can run in the browser.
	let { js, css } = svelte.compile(component, {
		css: false
	});
	  
	// Write JS to build directory.
	ensureDirExists(arg.destPath);
	fs.promises.writeFile(arg.destPath, js.code);

	// Write CSS to build directory.
	ensureDirExists(arg.stylePath);
	if (css.code && css.code != 'null') {
		fs.appendFileSync(arg.stylePath, css.code);
	}
});

// ------------------------
// Start static HTML build:
// ------------------------

let staticBuildStr = JSON.parse(args[1]);
let allNodes = JSON.parse(args[2]);

// Create the component that wraps all nodes.
let htmlWrapper = path.join(path.resolve(), 'layout/global/html.svelte')
const component = relative(htmlWrapper, process.cwd()).default;

staticBuildStr.forEach(arg => {

	let componentPath = path.join(path.resolve(), arg.componentPath);
	let destPath = path.join(path.resolve(), arg.destPath);

	// Set route used in svelte:component as "this" value.
	const route = relative(componentPath, process.cwd()).default;

	// Set props so component can access field values, etc.
	let props = {
		route: route,
		node: arg.node,
		allNodes: allNodes
	};

	// Create the static HTML and CSS.
	let { html, css } = component.render(props);

	// Inject Style.
	let style = "<style>" + css.code + "</style>";
	html = injectString('prepend', style, '</head>', html);
	// Inject SPA entry point.
	let entryPoint = '<script type="module" src="https://unpkg.com/dimport?module" data-main="/spa/ejected/main.js"></script><script nomodule src="https://unpkg.com/dimport/nomodule" data-main="/spa/ejected/main.js"></script>';
	html = injectString('prepend', entryPoint, '</head>', html);
	// Inject ID used to hydrate SPA.
	let hydrator = ' id="hydrate-plenti"';
	html = injectString('append', hydrator, '<html', html);

	// Write .html file to filesystem.
  	ensureDirExists(destPath);
	fs.promises.writeFile(destPath, html);
	  
});`),
	"/layout/ejected/main.js": []byte(`import Router from './router.svelte';

/*
if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/plenti-service-worker.js')
  .then((reg) => {
    console.log('Service Worker registration succeeded.');
  }).catch((error) => {
    console.log('Service Worker registration failed with ' + error);
  });
} else {
  console.log('Service Workers not supported by browser')
}
*/

const replaceContainer = function ( Component, options ) {
  const frag = document.createDocumentFragment();
  const component = new Component( Object.assign( {}, options, { target: frag } ));
  if (options.target) {
    options.target.replaceWith( frag );
  }
  return component;
}

const app = replaceContainer( Router, {
  target: document.querySelector( '#hydrate-plenti' ),
  props: {}
});

export default app;
`),
	"/layout/ejected/router.svelte": []byte(`<Html {route} {node} {allNodes} />

<script>
  import Navaid from 'navaid';
  import nodes from './nodes.js';
  import Html from '../global/html.svelte';

  let route, node, allNodes;

  const getNode = (uri, trailingSlash = "") => {
    return nodes.find(node => node.path + trailingSlash == uri);
  }

  let uri = location.pathname;
  node = getNode(uri);
  if (node === undefined) {
    node = getNode(uri, "/");
  }
  allNodes = nodes;

  function draw(m) {
    node = getNode(uri);
    if (node === undefined) {
      // Check if there is a 404 data source.
      node = getNode("/404");
      if (node === undefined) {
        // If no 404.json data source exists, pass placeholder values.
        node = {
          "path": "/404",
          "type": "404",
          "filename": "404.json",
          "fields": {}
        }
      }
    }
    route = m.default;
    window.scrollTo(0, 0);
  }

  function track(obj) {
    uri = obj.state || obj.uri;
  }

  addEventListener('replacestate', track);
  addEventListener('pushstate', track);
  addEventListener('popstate', track);

  const router = Navaid('/', () => {
    import('../content/404.js')
      .then(draw)
      .catch(err => {
        console.log("Add a '/layout/content/404.svelte' file to handle Page Not Found errors.");
        console.log("If you want to pass data to your 404 component, you can also add a '/content/404.json' file.");
        console.log(err);                                                                                           
      });
  });

  allNodes.forEach(node => {
    router.on(node.path, () => {
      // Check if the url visited ends in a trailing slash (besides the homepage).
      if (uri.length > 1 && uri.slice(-1) == "/") {
        // Redirect to the same path without the trailing slash.
        router.route(node.path, false);
      } else {
        import('../content/' + node.type + '.js').then(draw);
      }
    });

  });

  router.listen();

</script>
`),
	"/layout/global/footer.svelte": []byte(`<script>
  export let allNodes;
  import { makeTitle } from '../scripts/make_title.svelte';
</script>
<footer>
  <div class="container">
    <span>All nodes:</span>
    {#each allNodes as node}
      <a href="{node.path}">{makeTitle(node.filename)}</a>&nbsp;
    {/each}
  </div>
</footer>

<style>
  footer {
    min-height: 200px;
    display: flex;
    align-items: center;
    background-color: var(--base);
    margin-top: 100px;
  }
</style>
`),
	"/layout/global/head.svelte": []byte(`<script>
  export let title;
</script>

<head>
  <meta charset='utf-8'>
  <meta name='viewport' content='width=device-width,initial-scale=1'>

  <title>{ title }</title>

  <link href="https://fonts.googleapis.com/css2?family=Rubik:ital,wght@0,300;0,700;1,300&display=swap" rel="stylesheet">
  <link rel='icon' type='image/png' href='/favicon.png'>
  <link rel='stylesheet' href='/spa/bundle.css'>
</head>
`),
	"/layout/global/html.svelte": []byte(`<script>
  import Head from './head.svelte';
  import Nav from './nav.svelte';
  import Footer from './footer.svelte';
  import { makeTitle } from '../scripts/make_title.svelte';

  export let route, node, allNodes;
</script>

<html lang="en">
<Head title={makeTitle(node.filename)} />
<body>
  <Nav />
  <main>
    <div class="container">
      <svelte:component this={route} {...node.fields} {allNodes} />
      <br />
    </div>
  </main>
  <Footer {allNodes} />
</body>
</html>

<style>
  body {
    font-family: 'Rubik', sans-serif;
    display: flex;
    flex-direction: column;
    margin: 0;
  }
  main {
    flex-grow: 1;
  }
  :global(.container) {
    max-width: 1024px;
    margin: 0 auto;
    flex-grow: 1;
  }
  :global(:root) {
    --primary: (50, 50, 50);
    --accent: rgb(1, 1, 1);
    --base: rgb(245, 245, 245);
  }
</style>
`),
	"/layout/global/nav.svelte": []byte(`<nav>
  <div class="container">
    <span id="brand"><a href="/">Home</a></span>
    <a href="/about">About</a>&nbsp;
    <a href="/contact">Contact</a>
  </div>
</nav>

<style>
  nav {
    min-height: 60px;
    display: flex;
    align-items: center;
    box-shadow: 0px 2px 3px var(--base);
  }
  .container {
    display: flex;
  }
  #brand {
    flex: 1;
  }
</style>
`),
	"/layout/scripts/make_title.svelte": []byte(`<script context="module">
  export const makeTitle = filename => {
  if (filename == '_index.json') {
    return 'Home';
  } else if (filename) {
    // Remove file extension.
    filename = filename.split('.').slice(0, -1).join('.');
    // Convert underscores and hyphens to spaces.
    filename = filename.replace(/_|-/g,' ');
    // Capitalize first letter of each word.
    filename = filename.split(' ').map((s) => s.charAt(0).toUpperCase() + s.substring(1)).join(' ');
  } 
  return filename;
  }
</script>
`),
	"/layout/scripts/sort_by_date.svelte": []byte(`<script context="module">
  export const sortByDate = (items, order) => {
    items.sort((a, b) => { 
      // Must have a field specifically named "date" to work.
      // Feel free to extend to other custom named date fields.
      if (a.fields.hasOwnProperty("date") && b.fields.hasOwnProperty("date")) {
        let aDate = new Date(a.fields.date);
        let bDate = new Date(b.fields.date);
        if (order == "oldest") {
            return aDate - bDate;
        }
        return bDate - aDate;
      }
    });
    return items;
  };
</script>`),
	"/package.json": []byte(`{
  "name": "my-plenti-app",
  "version": "1.0.0",
  "type": "module",
  "private": true,
  "devDependencies": {
    "require-relative": "^0.8.7"
  },
  "dependencies": {
    "navaid": "^1.0.5",
    "regexparam": "^1.3.0",
    "svelte": "^3.21.0"
  }
}
`),
	"/plenti.json": []byte(`{
	"baseurl": "http://example.org/",
	"title": "My New Plenti Site",
	"types": {
		"pages": "/:filename"
	},
	"build": "public",
	"local": {
		"port": 3000
	}
}`),
}
