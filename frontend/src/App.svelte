<script lang="ts">
	import Router from "svelte-spa-router";
	import routes from "./routes";

	import "@fontsource/fira-sans";
	import "@fontsource/arvo";
	import "@fontsource/fira-code";

</script>

<main>
	<div class="bar">
		<div class="heading">
			<h1><a href="/#/">paste</a></h1>
		</div>
		<div class="links">
			<a href="/#/about">about</a>
			<!--<a href="/#/debug">debug</a>-->
		</div>
	</div>
	<div class="content">
		<Router {routes} />
	</div>
</main>

<style lang="scss">
	$color1: #edf2fb;
	$color2: #d7e3fc;
	$color3: #013a63;

	:global(body) {
		font-family: "Fira Sans";
	}

	:global(.aa-DropZone) {
		color: darkgrey;
	}

	main {
		max-width: 960px;
		min-width: 250px;
		margin: auto;
		padding: 0rem 2rem;

		a {
			color: $color3;
		}

		.bar {
			display: flex;
			justify-content: space-between;
			align-items: center;

			h1 > a {
				color: black;
				&:hover {
					text-decoration: none;
				}
			}
		}
	}
</style>
