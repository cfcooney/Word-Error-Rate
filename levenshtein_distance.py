import numpy as np 

def word_error_rate(str1, str2):
	"""
	Simple implementation of the Levenstein Distence and word error rate (WER)
	Distance computed with iterative matrix
	WER computed using formula from [2]
	[1] Levenshtein, Vladimir I. (February 1966). "Binary codes capable of correcting 
	    deletions, insertions, and reversals". Soviet Physics Doklady. 10 (8): 707–710.
	[2] Soukoreff, R. W., & MacKenzie, I. S. (2001). Measuring errors in text entry tasks:
	    An application of the Levenshtein string distance statistic. Extended Abstracts of the 
	    ACM Conference on Human Factors in Computing Systems - CHI 2001, pp. 319-320
		(WER: https://www.yorku.ca/mack/CHI01a.html)

	Python implementation of the levenshtein distance computation:
	input: str1: string
		   str2: string
    output: int
	"""
	if str1 == str2:
		return 0 
	else:

		d = np.zeros([len(str1)+1, len(str2)+1], dtype=int)

		for i in range(1, len(str1)+1):
			d[i, 0] = i

		for j in range(1, len(str2)+1):
			d[0, j] = j

		for j in range(1, len(str2)+1):
			for i in range(1, len(str1)+1):


				if str1[i-1] == str2[j-1]:
					substitution_cost = 0
				else:
					substitution_cost = 1

				
				updates = dict(a=d[i-1, j] + 1,b=d[i, j-1] + 1,c=d[i-1, j-1] + substitution_cost)
				key_min = min(updates.keys(), key=(lambda k: updates[k]))
				d[i, j] = updates[key_min]
				

		WER = (d[len(str1), len(str2)] / max(len(str1), len(str2))) * 100
		WAcc = 100 - WER
		return d[len(str1), len(str2)], WER, WAcc

distance, WER, WAcc = word_error_rate('hello','world')

print(f"Levenstein distance: {distance}")
print(f"Word error rate: {WER}%")
print(F"Word Accuracy: {WAcc}%")

	