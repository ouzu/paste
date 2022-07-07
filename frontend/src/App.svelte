<script lang="ts">
	import Router from "svelte-spa-router";
	import routes from "./routes";

	import { _, setupI18n } from "./i18n";

	import "@fontsource/fira-sans";
	import "@fontsource/fira-code";

	//setupI18n({ withLocale: "de" });
	setupI18n();
</script>

<main>
	<div class="bar">
		<div class="heading">
			<h1><a href="/#/">{$_("app.title")}</a></h1>
		</div>
		<div class="links">
			<a href="/#/about">{$_("app.aboutLink")}</a>
			<!-- client releases broken :(
			•
			<a href="/#/client">{$_("app.clientLink")}</a>
			-->
			<!--<a href="/#/debug">debug</a>-->
		</div>
	</div>
	<div class="content">
		<Router {routes} />
	</div>
</main>

<style lang="scss">
	:global(body) {
		font-family: "Fira Sans";
	}

	:global(.aa-DropZone) {
		color: var(--fg-color2);
		background-color: var(--bg-color);
	}

	main {
		max-width: 960px;
		min-width: 250px;
		margin: auto;
		padding: 0rem 2rem;

		.bar {
			display: flex;
			justify-content: space-between;
			align-items: center;

			h1 > a {
				color: var(--fg-color1);
				&:hover {
					text-decoration: none;
				}
			}
		}
	}
</style>
