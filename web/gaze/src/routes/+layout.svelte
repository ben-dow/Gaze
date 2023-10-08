<script>
    import "../app.css";
    import Navigation from "../lib/nav/Navigation.svelte";
</script>


<div class="flex h-screen">
    <div class="flex-1 flex flex-col overflow-hidden">
        <header class="flex justify-between items-center bg-slate-700 h-14 p-4">
            <h1 class="text-5xl font-mono text-slate-200">Gaze</h1>
        </header>
        <div class="flex h-full">
            <Navigation/>
            <main class="flex flex-col w-full  overflow-x-hidden overflow-y-auto shadow-inner shadow-slate-900 p-2">
               <slot />
            </main>
        </div>
    </div>
</div>


<style lang="postcss">
    :global(html) {
        background-color: theme(colors.slate.800);
    }

    /* width */
    ::-webkit-scrollbar {
        width: 5px;
    }

    /* Track */
    ::-webkit-scrollbar-track {
        background: #f1f1f1;
        border-radius: 100px;
    }

    /* Handle */
    ::-webkit-scrollbar-thumb {
        background: #888;
        border-radius: 100px;
    }

    /* Handle on hover */
    ::-webkit-scrollbar-thumb:hover {
        background: #555;
    }
</style>