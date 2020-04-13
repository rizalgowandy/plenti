package generated

// Do not edit, this file is automatically generated.

// Defaults: scaffolding used in 'new site' command
var Defaults = map[string][]byte{
	"/.gitignore": []byte(`public
node_modules`),
	"/content/index.json": []byte(`{
	"title": "My Site Homepage",
	"intro": {
		"slogan": "Welcome to a faster way to web",
		"color": "red"
	}
}`),
	"/content/pages/_blueprint.json": []byte(`{
	"title": "",
	"desc": "",
	"author": ""
}`),
	"/content/pages/about.json": []byte(`{
	"title": "About Me",
	"desc": "Tell us about yourself",
	"author": "Your name"
}`),
	"/content/pages/contact.json": []byte(`{
	"title": "Contact",
	"desc": "Maybe add a <a href='https://plentiform.com'>plentiform</a>?",
	"author": "Your name"
}`),
	"/layout/components/grid.svelte": []byte(`<script>
  export let items, filter;
</script>

<div class="grid">
  {#each items as item}
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
	"/layout/content/blog.svelte": []byte(`<script>
	export let title, description;
</script>

<h1>{title}</h1>
<p><em>Blog template</em></p>
<div>
  <div><strong>Title: </strong><span>{title}</span></div>
  <div><strong>Desc: </strong><span>{description}</span></div>
</div>

<p><a href="/">Back home</a></p>
`),
	"/layout/content/index.svelte": []byte(`<script>
	export let name;
	export let allNodes;
  import Grid from '../components/grid.svelte';
</script>

<h1>{name}</h1>
<section id="intro">
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<h3>Recent blog posts:</h3>
  <Grid items={allNodes} filter="blog" />
</section>
`),
	"/layout/content/pages.svelte": []byte(`<script>
	export let title, description;
</script>

<h1>{title}</h1>
<p><em>Page template</em></p>
<div>
  <div><strong>Title: </strong><span>{title}</span></div>
  <div><strong>Desc: </strong><span>{description}</span></div>
</div>

<p><a href="/">Back home</a></p>
`),
	"/layout/ejected/client_router.svelte": []byte(`<Html {Route} {node} {allNodes} />

<script>
  import Navaid from 'navaid';
  import DataSource from './data_source.js';
  import { onDestroy } from 'svelte';
  import Html from '../global/html.svelte';

  let Route, node, allNodes;

  let uri = location.pathname;
  node = DataSource.getNode(uri);
  allNodes = DataSource.getAllNodes();

  function draw(m) {
    Route = m.default;
    window.scrollTo(0, 0);
  }

  function track(obj) {
    uri = obj.state || obj.uri;
    if (window.ga) ga.send('pageview', { dp:uri });

    node = DataSource.getNode(uri);
    allNodes = DataSource.getAllNodes();
  }

  addEventListener('replacestate', track);
  addEventListener('pushstate', track);
  addEventListener('popstate', track);

  const router = Navaid('/')
    .on('/', () => import('../content/index.svelte').then(draw))
    .on('/:slug', () => import('../content/pages.svelte').then(draw))
    .on('/blog/:slug', () => import('../content/blog.svelte').then(draw))
    .listen();

  onDestroy(router.unlisten);
</script>
`),
	"/layout/ejected/data_source.js": []byte(`import nodes from './nodes.js';

class DataSource {

  constructor() {}

  static getNode(uri) {
    let content;
    nodes.map(node => {
      if (node.path == uri) {
        content = node;
      }
    });
    return content ? content : '';
  }

  static getAllNodes() {
    let content = nodes.map(node => {
      return node;
    });
    return content;
  }
}

export default DataSource;
`),
	"/layout/ejected/main.js": []byte(`import Router from './client_router.svelte';

if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/jim-service-worker.js')
  .then((reg) => {
    console.log('Service Worker registration succeeded.');
  }).catch((error) => {
    console.log('Service Worker registration failed with ' + error);
  });
} else {
  console.log('Service Workers not supported by browser')
}

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
	"/layout/ejected/nodes.js": []byte(`// TODO: This file is temporary and needs to be removed at some point.
// It should get automatically generated by plenti based on the "content/" folder.
const nodes = [
    {
        "path": "/blog/post1",
        "type": "blog",
        "filename": "post1.json",
        "fields": {
            "title": "Post 1",
            "description": "First blog post."
        }
    },
    {
        "path": "/blog/post2",
        "type": "blog",
        "filename": "post2.json",
        "fields": {
            "title": "Post 2",
            "description": "Second blog post."
        }
    },
    {
        "path": "/blog/post3",
        "type": "blog",
        "filename": "post-3_has_a_long_filename.json",
        "fields": {
            "title": "Post 3",
            "description": "Third of the blog posts."
        }
    },
    {
        "path": "/about",
        "type": "pages",
        "filename": "about.json",
        "fields": {
            "title": "About Page",
            "description": "This is the about page"
        }
    },
    {
        "path": "/anything",
        "type": "pages",
        "filename": "anything.json",
        "fields": {
            "title": "Anything!",
            "description": "The amazing anything page..."
        }
    },
    {
        "path": "/",
        "type": "index",
        "filename": "index.json",
        "fields": {
            "name": "Plenti"
        }
    }
];

export default nodes;
`),
	"/layout/ejected/server_router.js": []byte(`import path from 'path';
import fs from 'fs';
import nodes from './nodes.js';
import 'svelte/register.js';
import relative from 'require-relative';

const injectString = (order, content, element, html) => {
	if (order == 'prepend') {
		return html.replace(element, content + element);
	} else if (order == 'append') {
		return html.replace(element, element + content);
	}
};

const ensureDirExists = filePath => {
	var dirname = path.dirname(filePath);
	if (fs.existsSync(dirname)) {
		return true;
	}
	ensureDirExists(dirname);
	fs.mkdirSync(dirname);
}

nodes.forEach(node => {
  let sourcePath = path.join(path.resolve(), 'layout/content/' + node.type + '.svelte');
  let sourceComponent = fs.readFileSync(sourcePath, 'utf8');
  let index = node.filename == 'index.json' ? 'index' : '';
  let destPath = path.join(path.resolve(), 'public/' + node.path + index + ".html");
  let topLevelComponent = path.join(path.resolve(), 'layout/global/html.svelte');
  const route = relative(sourcePath, process.cwd()).default;
  let props = {
    Route: route,
    node: node,
    allNodes: nodes
  };
  // Create HTML file.
  const component = relative(topLevelComponent, process.cwd()).default;
  let { html, css } = component.render(props);
  // Inject Style.
  let style = "<style>" + css.code + "</style>";
  html = injectString('prepend', style, '</head>', html);
  // Inject SPA entry point.
  let entryPoint = '<script type="module" src="https://unpkg.com/dimport?module" data-main="/.spa/main.js"></script><script nomodule src="https://unpkg.com/dimport/nomodule" data-main="/.spa/main.js"></script>';
  html = injectString('prepend', entryPoint, '</head>', html);
  // Inject ID used to hydrate SPA.
  let hydrator = ' id="hydrate-plenti"';
  html = injectString('append', hydrator, '<html', html);
  // Write HTML files to filesystem.
  ensureDirExists(destPath);
  fs.promises.writeFile(destPath, html);
});
`),
	"/layout/global/footer.svelte": []byte(`<script>
  export let allNodes;
</script>
<footer>
  <div class="container">
    <span>All nodes:</span>
    {#each allNodes as node}
      {#if node.fields.title}
        <a href="{node.path}">{node.fields.title}</a>&nbsp;
      {/if}
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
  <link rel='stylesheet' href='/.spa/bundle.css'>
</head>
`),
	"/layout/global/html.svelte": []byte(`<script>
  import Head from './head.svelte';
  import Nav from './nav.svelte';
  import Footer from './footer.svelte';
  import { makeTitle } from '../scripts/make_title.svelte';

  export let Route, node, allNodes;
</script>

<html lang="en">
<Head title={makeTitle(node.filename)} />
<body>
  <Nav />
  <main>
    <div class="container">
      <svelte:component this={Route} {...node.fields} {allNodes} />
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
    <a href="/anything">Anything</a>
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
	"/package.json": []byte(`{
  "name": "svelte-app",
  "version": "1.0.0",
  "type": "module",
  "scripts": {
    "build": "rollup -c",
    "dev": "rollup -c -w",
    "start": "sirv public"
  },
  "devDependencies": {
    "@rollup/plugin-commonjs": "^11.0.0",
    "@rollup/plugin-node-resolve": "^6.0.0",
    "rollup": "^1.20.0",
    "rollup-plugin-livereload": "^1.0.0",
    "rollup-plugin-svelte": "^5.0.3",
    "rollup-plugin-terser": "^5.1.2",
    "sirv-cli": "^0.4.4",
    "snowpack": "^1.6.0",
    "svelte": "^3.0.0"
  },
  "dependencies": {
    "navaid": "^1.0.5",
    "require-relative": "^0.8.7"
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
	"/rollup.config.js": []byte(`import svelte from 'rollup-plugin-svelte';
import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';

const production = !process.env.ROLLUP_WATCH;

export default [{
    input: 'layout/ejected/main.js',
    output: {
        sourcemap: true,
        format: 'esm',
        name: 'app',
        dir: 'public/.spa'
    },
    plugins: [
        svelte({
            dev: !production,
            css: css => {
                css.write('public/.spa/bundle.css');
            }
        }),
        resolve({
            browser: true,
            dedupe: importee => importee === 'svelte' || importee.startsWith('svelte/')
        }),
        commonjs(),
        !production && serve(),
        !production && livereload('public'),
        production && terser()
    ],
    watch: {
        clearScreen: false
    }
}];

function serve() {
    let started = false;

    return {
        writeBundle() {
            if (!started) {
                started = true;

                require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
                    stdio: ['ignore', 'inherit', 'inherit'],
                    shell: true
                });
            }
        }
    };
}
`),
}