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
        equipment_list = await List()
        equipment_type_list = await ListEquipmentTypes()
        usr = await LoadUser()
    })
</script>
{#if usr}
    <EquipmentViewer bind:usr={usr} bind:equipment_list={equipment_list} bind:equipment_type_list={equipment_type_list} />
{:else}
    Loading EquipmentViewer...
{/if}