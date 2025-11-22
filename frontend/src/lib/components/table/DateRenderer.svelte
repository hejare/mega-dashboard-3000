<script lang="ts">
    export let date: string;
</script>

<div>
    {new Date(date).toLocaleDateString("sv-SE")}
</div>