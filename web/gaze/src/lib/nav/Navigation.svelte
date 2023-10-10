<script>
    import ExpandableNavItem from "$lib/nav/ExpandableNavItem.svelte";
    import NavItem from "./NavItem.svelte";
</script>

<nav class="flex w-72 h-full bg-gray-900 p-4 flex-col justify-start text-slate-200 gap-y-4 overflow-auto border-r-2 border-r-gray-600 rounded">
    <ExpandableNavItem itemName="Libraries" expand={false}/>
    <ExpandableNavItem itemName="Collections" expand={false}/>
    <NavItem itemName="Import"/>
    <NavItem itemName="Settings"/>
</nav>