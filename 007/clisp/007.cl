#!/bin/bash
#|
exec clisp -q -q $0 $0 ${1+"$@"}
exit
|#
(setf lista '(2 3 5 7 11 13))
(setf num 15)  
(setf fim 10001)

(defun to(num)
  (setf conta 0)
  (loop for x in lista
      do (when (> (mod num x) 0)
	(setf conta (1+ conta)))
      do (when (= conta (length lista))
	  (nconc lista (list num)))))

(defun main()
(loop while (< 2 num)
  do (to num) 
  do (when (= fim (length lista))
   (print (last lista))
   (quit))
  do (setf num (1+ num))))  
  
(time (main))