from Infer import Infer

params = {
    'BATCH_SIZE' : 1,
    'PLOT_TOP_K': 10,
    'INFER_PATH': './Data/Imgs',
}
i = Infer(params)
i.infer()