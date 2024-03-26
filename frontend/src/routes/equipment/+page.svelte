<script lang="ts">
    import { onMount } from 'svelte';
    import EquipmentViewer from '$lib/components/equipment/EquipmentViewer.svelte';
    import type { model } from '$lib/wailsjs/go/models';
    import { List, ListEquipmentTypes } from '$lib/wailsjs/go/controllers/EquipmentHandler';
    import { LoadUser } from '$lib/wailsjs/go/main/App';

    let equipment_list: model.Equipment[] = [];
    let equipment_type_list: model.EquipmentType[] = [];
    let usr: model.User;
    onMount(async () => {
        usr = await LoadUser()
        equipment_list = await List()
        equipment_type_list = await ListEquipmentTypes()
    })
</script>
{#if equipment_type_list.length > 0 && usr}
    <EquipmentViewer bind:usr={usr} bind:equipment_list={equipment_list} equipment_type_list={equipment_type_list}></EquipmentViewer>
{/if}