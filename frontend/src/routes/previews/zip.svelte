<script>
    export let file;
    export let name;

    import { onMount } from "svelte";

    import JSZip from "jszip";

    var zip = new JSZip();

    let files = [];

    onMount(async () => {
        await zip.loadAsync(file);
        files = Object.keys(zip.files);
        files.sort();
    });
</script>

<div class="zip">
    <p>archive contents:</p>
    <ul>
        {#each files as f}
            <li>{f}</li>
        {/each}
    </ul>
</div>

<style lang="scss">
    ul { 
        list-style: none;
    }
</style>
