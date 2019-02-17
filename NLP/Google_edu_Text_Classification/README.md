# Text Classification

*Training link:*  
https://developers.google.com/machine-learning/guides/text-classification/?authuser=0  
*The code is from:*  
https://github.com/google/eng-edu/blob/master/ml/guides/text_classification/README.md  

This repository contains end-to-end tutorial-like code samples to help solve
text classification problems using machine learning.

## Modules

We have one module for each step in the text classification workflow.

1.  *load_data* - Functions to load data from four different datasets. For each
    of the dataset we:

    +   Read the required fields (texts and labels).
    +   Do any pre-processing if required. For example, make sure all label
        values are in range [0, num_classes-1].
    +   Split the data into training and validation sets.
    +   Shuffle the training data.

2.  *explore_data* - Helper functions to understand datasets.

3.  *vectorize_data* - N-gram and sequence vectorization functions.

4.  *build_model* - Helper functions to create multi-layer perceptron and
    separable convnet models.

5.  *train data* - Demonstrates how to use all of the above modules and train a
    model.

    + *train_ngram_model* - Trains a multi-layer perceptron model on IMDb
    movie reviews sentiment analysis dataset.

    + *train_sequence_model* - Trains a sepCNN model on Rotten Tomatoes movie
    reviews sentiment analysis dataset.

    + *train_fine_tuned_sequence_model* - Trains a sepCNN model with
    pre-trained embeddings that are fine-tuned on Tweet weather topic
    classification dataset.

    + *batch_train_sequence_model* - Same as *train_sequence_model* but here
    we are training data in batches.

6.  *tune_model* - Contains example to demonstrate how you can find the best
    hyper-parameter values for your model.

## Algorithm for Data Preparation and Model Building:

![Alt text](img/TextClassificationFlowchart.png?raw=true "TextClassificationFlowchart")

1. Calculate the number of samples/number of words per sample ratio.
2. If this ratio is less than 1500, tokenize the text as n-grams and use a
simple multi-layer perceptron (MLP) model to classify them (left branch in the
flowchart below):
  a. Split the samples into word n-grams; convert the n-grams into vectors.
  b. Score the importance of the vectors and then select the top 20K using the scores.
  c. Build an MLP model.
3. If the ratio is greater than 1500, tokenize the text as sequences and use a
   sepCNN model to classify them (right branch in the flowchart below):
  a. Split the samples into words; select the top 20K words based on their frequency.
  b. Convert the samples into word sequence vectors.
  c. If the original number of samples/number of words per sample ratio is less
     than 15K, using a fine-tuned pre-trained embedding with the sepCNN
     model will likely provide the best results.
4. Measure the model performance with different hyperparameter values to find
   the best model configuration for the dataset.
   
