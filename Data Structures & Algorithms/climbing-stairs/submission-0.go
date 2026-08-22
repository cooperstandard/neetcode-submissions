func climbStairs(n int) int {
   // at most n unique paths

   // at each index keep track of the number of steps hitting that index and the number of steps hitting the next index
   // the number of paths is the number of steps which hit the last index

   this, next := 1,0

   for i := 0; i < n; i ++ {
    swap := this
    this = this + next
    next = swap
   }


   return this
}
