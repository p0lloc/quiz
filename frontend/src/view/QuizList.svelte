<script lang="ts">
    import { onMount } from "svelte";
    import QuizCard from "../components/admin/QuizCard.svelte";
    import type { Quiz } from "../model/quiz";
    import QuizQuestionCard from "../components/admin/QuizQuestionCard.svelte";
    import CreateQuiz from "../components/admin/CreateQuiz.svelte";

    let quizzes: Quiz[] = [
        {
            id: "math",
            name: "Math quiz",
            questions: [
                {
                    id: crypto.randomUUID(),
                    name: "What is 2+2?",

                    choices: [
                        {
                            id: crypto.randomUUID(),
                            value: "4",
                            correct: true,
                        },
                        {
                            id: crypto.randomUUID(),
                            value: "42",
                            correct: false,
                        },
                        {
                            id: crypto.randomUUID(),
                            value: "cat",
                            correct: false,
                        },
                        {
                            id: crypto.randomUUID(),
                            value: "π",
                            correct: false,
                        }
                    ]
                }
            ]
        },
        {
            id: crypto.randomUUID(),
            name: "Dummy quiz",
            questions: []
        }
    ];

    async function createQuiz(e: CustomEvent<Quiz>){
        const quiz = e.detail;
        let response = await fetch(`http://localhost:3000/api/quizzes`, {
            method: "post",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                name: quiz.name,
                questions: quiz.questions
            })
        });

        let json = await response.json();
        quizzes = json;
    }

    async function getQuizzes(){
        let response = await fetch(`http://localhost:3000/api/quizzes`);
        let json = await response.json();
        quizzes = json;
    }

    onMount(async () => {
        await getQuizzes();
    });
</script>

<div class="p-4">
    <div class="flex flex-col gap-2">
        {#each quizzes as quiz, i (quiz.id) }
            <QuizCard {quiz} />
        {/each}
    </div>

    <div class="mt-10">
        <CreateQuiz on:create={createQuiz} />
    </div>
</div>