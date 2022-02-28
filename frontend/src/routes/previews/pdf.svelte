<script>
    export let url;
    export let name;

    import { _ } from "../../i18n";

</script>

<div class="pdf">
    <object title={name} class="display-pdf" data={url} type="application/pdf">
        {$_("file.pdf.warning")}
    </object>
</div>

<style lang="scss">
    .pdf {
        margin: 0 -1rem 1rem;

        .display-pdf {
            width: 100%;
            aspect-ratio: 1 / 1;
            border-radius: 0.3rem;
            //box-shadow: 0 10px 20px rgba(0,0,0,0.19), 0 6px 6px rgba(0,0,0,0.23);
        }
    }
</style>
