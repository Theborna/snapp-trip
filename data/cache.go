package data

/*
Caching implementation:

we will be using redis for all caching purposes in this project.
for our purposes we will be caching the most useful names in our database to reduce the number of
calls we make to our database.
*/

const CACHE_SIZE = 1000
