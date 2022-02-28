<script>
    export let file;

    import { onMount } from "svelte";

    import { _ } from "../../i18n";
    
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
    <p>{$_("file.zip.contents")}</p>
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
