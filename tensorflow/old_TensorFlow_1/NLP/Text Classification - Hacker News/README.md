# Text Classification for Hacker News

Hacker news headlines are available as a BigQuery public dataset.   
The dataset contains all headlines from the sites inception in October 2006 until October 2015.  

Creating datasets for Machine Learning using BigQuery ( already exported in tsv files)  
Creating a text classification model using the Estimator API with a Keras model  

The steps:  

+ tf.keras.preprocessing.text.Tokenizer.fit_on_texts() to generate a mapping from our word vocabulary to integers 
+ tf.keras.preprocessing.text.Tokenizer.texts_to_sequences() to encode our sentences into a sequence of their respective word-integers 
+ tf.keras.preprocessing.sequence.pad_sequences() to pad all sequences to be the same length 
+ The embedding layer in the keras model takes care of one-hot encoding these integers and learning a dense emedding represetation from them.

Finally we pass the embedded text representation through a CNN model pictured below

![Alt text](images/txtcls_model.png?raw=true "Model Flow")
