<script lang="ts">
    import type { QuizChoice } from "../../model/quiz";

    export let choice: QuizChoice;
</script>

<div class="flex gap-2">
    <input type="text" placeholder="Choice" bind:value={choice.value}>
    <input type="checkbox" bind:checked={choice.correct}>
</div>