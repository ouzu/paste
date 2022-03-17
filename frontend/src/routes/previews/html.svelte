<script>
    export let text;
    export let name;
</script>

<div class="html">
    <!-- FIXME: NO XSS PROTECTION! -->
    {@html text}
</div>

<style lang="scss">
    .html {
        background-color: var(--ac-1);
        border-radius: 0.3rem;
        padding: 0rem 1rem;
        margin: 0 -1rem 1rem;
        overflow: hidden;
    }
</style>
