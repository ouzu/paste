<script>
    export let text;
    export let name;
    import Prism from "prismjs";
    import "prismjs/themes/prism.min.css";
    import { getLanguage } from "./helper";

    import "@fontsource/fira-code";

    let language = getLanguage(name);
</script>

<div class="code">
    <div class="inner">
        {@html Prism.highlight(text, Prism.languages[language], language)}
    </div>
</div>

<style lang="scss">
    .code {
        margin: 0 -1rem 1rem;
        white-space: pre;
        font-family: "Fira Code";
        overflow: auto;
        border-radius: 0.3rem;
        background-color: var(--ac-1);

        .inner {
            margin: 1rem 1rem;
        }
    }

    :global(.language-css .token.string, .style .token.string, .token.entity, .token.operator, .token.url) {
        background: transparent !important;
    }
</style>
