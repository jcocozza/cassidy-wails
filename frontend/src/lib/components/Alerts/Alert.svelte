<script lang="ts">
    import { RemoveAlert, type Alert } from "../../../stores/alert";

    export let alert: Alert;

    function AlertId(): number {
        return Math.floor(Math.random() * 10000)
    }
</script>


<div class="alert alert-danger" role="alert">
    {alert.Message}
    <button type="button" class="btn-close" on:click={() => {RemoveAlert(alert.Id)}}></button>
</div>