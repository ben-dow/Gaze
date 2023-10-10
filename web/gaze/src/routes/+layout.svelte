<script>
    import "../app.css";
    import Navigation from "../lib/nav/Navigation.svelte";
    import Header from "../lib/header/Header.svelte";
</script>


<div class="flex h-screen">
    <div class="flex-1 flex flex-col overflow-hidden">
        <Header/>
        <div class="flex h-full">
            <Navigation/>
            <main class="flex flex-col w-full  overflow-x-hidden overflow-y-auto  p-2">
               <slot />
            </main>
        </div>
    </div>
</div>


<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.950);
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