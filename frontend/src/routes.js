import Home from './routes/Home.svelte';
import Lorem from './routes/Lorem.svelte';
import NotFound from './routes/NotFound.svelte';
import About from './routes/About.svelte';
import Debug from './routes/Debug.svelte';
import File from './routes/File.svelte';

// import type {RouteDefinition} from 'svelte-spa-router';

// const routes: RouteDefinition = {
const routes = {
    '/': Home,
    '/about': About,
    '/debug': Debug,
    '/lorem/:repeat': Lorem,
    '/files/:iv/:key': File,
    // The catch-all route must always be last
    '*': NotFound
};

export default routes;
