;;; anagram.el --- Anagram (exercism)

;;; Commentary:

;;; Code:

(require 'cl-lib)
(require 'seq)

(defun freq (w)
  (let ((mht (make-hash-table)))
    (progn (mapcar (lambda (c) (puthash c (+ 1 (gethash c mht 0)) mht)) w)
	   mht)))

 (defun hash-to-list (hash-table)
   (let (result-my)
     (maphash
      (lambda (k v)
	(push (list k v) result-my))
      hash-table)
     result-my))

(defun map-equal (m1 m2)
  (seq-every-p 'identity (mapcar (lambda (x) (equal (gethash (car x) m2 0) (cadr x))) (hash-to-list m1))))


(defun anagrams-for (w ws)
  (let ((w-mht (freq (downcase w))))
    (seq-filter (lambda (ow) (and (not (string= (downcase w) (downcase ow))) (eq (length ow) (length w)) (map-equal w-mht (freq (downcase ow))))) ws)))

(provide 'anagram)
;;; anagram.el ends here

