<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import type { Quiz } from "../../model/quiz";
    import QuizQuestionCard from "./QuizQuestionCard.svelte";

    const dispatch = createEventDispatcher();

    let quiz: Quiz = {
        id: "",
        name: "",
        questions: [
            {
                id: crypto.randomUUID(),
                name: "",
                choices: [
                    {
                        id: crypto.randomUUID(),
                        value: "",
                        correct: false
                    },
                    {
                        id: crypto.randomUUID(),
                        value: "",
                        correct: false
                    },
                    {
                        id: crypto.randomUUID(),
                        value: "",
                        correct: true
                    },
                    {
                        id: crypto.randomUUID(),
                        value: "",
                        correct: false
                    }
                ]
            }
        ]
    }

    function createQuiz(){
        dispatch("create", quiz);
    }
    
</script>

<h2 class="text-2xl font-bold mb-2">Create quiz</h2>
<input type="text" placeholder="Quiz name" bind:value={quiz.name} />
<div class="mb-4 mt-6">
    {#each quiz.questions as question(question.id)}
        <QuizQuestionCard {question} />
    {/each}
</div>
<button on:click={createQuiz}>Create quiz</button>