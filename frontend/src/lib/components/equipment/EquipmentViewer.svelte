<script lang="ts">
    import { onMount } from "svelte";
    import EquipmentCard from "./EquipmentCard.svelte";
    import NewEquipmentModal from "./NewEquipmentModal.svelte";
    import type { model } from "../../wailsjs/go/models";
    import { List } from '../../wailsjs/go/controllers/EquipmentHandler'

    export let equipment_list: model.Equipment[];
    export let equipment_type_list: model.EquipmentType[];

    async function refreshEquipment() {
        equipment_list = await List()
    }

    onMount(() => {
        refreshEquipment(); // Fetch initial data when the component is mounted
    });
</script>

<div class="equipment-list container">
    <div class="row">
        {#each equipment_list as equipment}
            <div class="col-md-4 mb-3">
                <EquipmentCard
                    bind:equipment={equipment}
                    on:delete={async () => {await refreshEquipment()}}
                />
            </div>
        {/each}
    </div>

    <NewEquipmentModal
        bind:equipment_type_list={equipment_type_list}
        on:update={async () => {await refreshEquipment()}}
    />
</div>