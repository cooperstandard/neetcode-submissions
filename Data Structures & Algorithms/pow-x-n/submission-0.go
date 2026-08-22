func myPow(x float64, n int) float64 {
   if x == 0 {
    return 0
   }

   if n < 0 {
    return 1 / myPow(x, n * -1)
   } 

   if n == 0 {
    return 1
   }

   if n == 1 {
    return x
   }

   pow := myPow(x, n/2) 
   pow = pow * pow
   if n % 2 == 1 {
    pow = pow * x
   }
   return pow
}


