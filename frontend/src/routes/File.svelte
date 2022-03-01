<script>
    export let params = {};

    import { onMount } from "svelte";
    import filesize from "filesize";

    import Image from "./previews/image.svelte";
    import Pdf from "./previews/pdf.svelte";
    import Audio from "./previews/audio.svelte";
    import Video from "./previews/video.svelte";
    import Text from "./previews/text.svelte";
    import Zip from "./previews/zip.svelte";

    import { _ } from "../i18n";

    import ProgressBar from "svelte-progress-bar";

    import { getEncoding } from "istextorbinary";

    import {
        importKey,
        importIV,
        decrypt,
        decryptString,
    } from "../crypto/crypto";

    let file = new Blob([]);
    let url = "";
    let name = "";
    let type = "";
    let text = "";
    let size = "";

    let key;
    let iv;
    let resp;

    let metaReady = false;
    let previewReady = false;

    let progress;

    let binary;

    const decoder = new TextDecoder();

    async function loadFile() {
        const responseReader = resp.body.getReader();

        let received = 0;

        let chunks = [];

        while (true) {
            const { done, value } = await responseReader.read();

            if (done) break;

            chunks.push(value);
            received += value.length;

            progress.setWidthRatio(received / size);
        }

        let blob = new Blob(chunks);

        let fileReader = new FileReader();
        fileReader.onload = async (e) => {
            let cryptData = e.target.result;

            let data = await decrypt(cryptData, iv, key);

            binary = getEncoding(decoder.decode(data)) === "binary";

            if (!binary) {
                text = decoder.decode(data);
            }

            file = new Blob([data]);
            url = URL.createObjectURL(file);

            previewReady = true;
            progress.setWidthRatio(0);
        };
        fileReader.readAsArrayBuffer(blob);
    }

    onMount(async () => {
        key = await importKey(params.key);
        iv = importIV(params.iv);

        let meta = await (
            await fetch("/api/meta/" + params.iv)
        ).json();
        name = await decryptString(meta.Name, iv, key);
        type = await decryptString(meta.Type, iv, key);

        resp = await fetch("/api/files/" + params.iv);
        size = Number(resp.headers.get("content-length"));

        metaReady = true;

        loadFile();
    });
</script>

<div class="file">
    <ProgressBar bind:this={progress} color="#abc4ff" />
    {#if metaReady}
        <div class="container">
            <div class="bar">
                <p>{name}</p>
                <p>
                    <span>{filesize(size)}</span>
                    <span>
                        <a
                            href={url}
                            download={name}
                            class:disabled={!previewReady}
                        >
                            {$_("file.save")}
                        </a>
                    </span>
                </p>
            </div>
        </div>
    {:else}
        <div class="spinner">
            <!-- By Sam Herbert (@sherb), for everyone. More @ http://goo.gl/7AJzbL -->
            <svg
                width="38"
                height="38"
                viewBox="0 0 38 38"
                xmlns="http://www.w3.org/2000/svg"
                stroke="#000"
            >
                <g fill="none" fill-rule="evenodd">
                    <g transform="translate(1 1)" stroke-width="2">
                        <circle stroke-opacity=".5" cx="18" cy="18" r="18" />
                        <path d="M36 18c0-9.94-8.06-18-18-18" />
                    </g>
                </g>
            </svg>
        </div>
    {/if}

    {#if previewReady}
        {#if type.startsWith("image/")}
            <Image {url} {name} />
        {:else if type === "application/pdf"}
            <Pdf {url} {name} />
        {:else if type.startsWith("audio/")}
            <Audio {url} {name} />
        {:else if type.startsWith("video/")}
            <Video {url} {name} />
        {:else if type === "application/zip"}
            <div class="container zip">
                <Zip {file} />
            </div>
        {:else if !binary}
            <Text {text} {name} />
        {:else}
            <p class="message">File can't be previewed.</p>
        {/if}
    {:else if metaReady}
        <div class="progress" />
    {/if}
</div>

<style lang="scss">
    .spinner {
        text-align: center;
        padding-top: 5rem;
        svg {
            animation: spin 1.2s linear infinite;
        }
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

    .message {
        text-align: center;
        margin-top: 2rem;
        color: var(--fg-color2);
    }

    .container {
        background-color: var(--ac-1);
        border-radius: 0.3rem;
        padding: 0rem 1rem;
        margin: 0 -1rem 1rem;

        &.zip {
            padding-top: 0.2rem;
            padding-bottom: 0.2rem;
        }

        .bar {
            display: flex;
            justify-content: space-between;
            align-items: center;

            span {
                margin: 0.2rem;
            }

            a.disabled {
                pointer-events: none;
                color: #aaa;
            }
        }
    }
</style>
