<script>
    import { _ } from "../i18n";

    const releases = [
        { os: "linux", arch: ["amd64", "arm64", "arm5"] },
        { os: "darwin", arch: ["amd64", "arm64"] },
        { os: "windows", arch: ["amd64"], "ext": ".exe" },
        { os: "android", arch: ["arm64"] },
    ];
</script>

<div class="client">
    <h1>{$_("client.heading")}</h1>
    <div>
        <p>{$_("client.summary")}</p>
    </div>

    <h2>{$_("client.usage")}</h2>
    <div>
        <p>{$_("client.upload")}</p>
        <code>pst upload [file]</code>
        <p>{$_("client.download")}</p>
        <code>pst download [link]</code>
        <p>{$_("client.help")}</p>
        <code>pst help</code>
    </div>

    <h2>{$_("client.releases")}</h2>
        {#each releases as r}
            <div class="release">
                <p>{r["os"]}:</p>
                <ul>
                    {#each r["arch"] as a}
                        <li>
                            <a href="/releases/{r['os']}-{a}/pst{r["ext"] || ""}" download
                                >{a}</a
                            >
                        </li>
                    {/each}
                </ul>
            </div>
        {/each}
</div>

<style lang="scss">
    div.client {
        margin: 2rem 0;
        div {
            margin: 1rem;
            &.release {
                display: inline-block;
                vertical-align: top;
                margin: -1rem 1rem 0;
                ul {
                    padding: 0;
                    list-style: none;
                }
            }
        }
    }
</style>
