<template>
    <nav>
        <ul>
            <li>
                <RouterLink class="link" to="/">Home</RouterLink>
            </li>
            <li>
                <RouterLink class="link" to="/addEmployee">Add Employee</RouterLink>
            </li>
        </ul>
    </nav>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>

<style lang="scss" scoped>
nav {
    padding: 24px 18px;
    display: inline-block;
    border-radius: 14px;

    ul {
        list-style-type: none;
        display: flex;
        text-align: center;
        padding-left: 0;

        li {
            padding: 0 8px;

            .link {
                text-decoration: none;
                color: rgb(41, 169, 255);

                &:hover {
                    color: rgb(0, 61, 102);
                }
            }
        }
    }
}
</style>