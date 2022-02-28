import About from './routes/About.svelte';
import File from './routes/File.svelte';
import Home from './routes/Home.svelte';
import NotFound from './routes/NotFound.svelte';

const routes = {
    '/': Home,
    '/about': About,
    '/files/:iv/:key': File,
    // The catch-all route must always be last
    '*': NotFound
};

export default routes;
