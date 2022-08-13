How to build the code
=====================

1. Please unzip the file "PriceBasket_Solution_SaikatRay.zip"
2. Open any bash/'command line' tools and 'change directory' to the extracted location
3. Please check the go Version by running the below command
    "go version"
4. Check that the go version installed is 18.0 or higher
5. Run the below command to build the code -
    "go build ."

How to run the code
===================

6. Run the below command to run the code -
    "go run ."
7. The program will ask for user input 
   (For getting valid results please enter the input in the format specfifed in the assignment)


Some features of the Code not specifically mentioned in the assignment
======================================================================

1. The Code supports altering the currency. The developer just has to change a constant 
   value at one place indicating the higher and lower denominations (like £ and p respectively) and add the conversion rule

2. Each .go files represents an entity with responsibility segregation. The main function only calls 
   the diffferent entities without storing any business logic.

3. All user inputs are case InSensitive and validated.

4. Due to time constarint many of the places has been marked that could be affected by concurrency instead of 
   directly applying GoRoutines.

5. Code could be simplified more by adding more helper methods for a common type. 

