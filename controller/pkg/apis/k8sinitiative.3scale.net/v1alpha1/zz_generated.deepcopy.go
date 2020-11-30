package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Product) DeepCopyInto(out *Product) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *Product) DeepCopy() *Product {
	if in == nil {
		return nil
	}
	out := new(Product)
	in.DeepCopyInto(out)
	return out
}

func (in *Product) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ProductList) DeepCopyInto(out *ProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Product, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *ProductList) DeepCopy() *ProductList {
	if in == nil {
		return nil
	}
	out := new(ProductList)
	in.DeepCopyInto(out)
	return out
}

func (in *ProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *ProductSpec) DeepCopyInto(out *ProductSpec) {
  *out = *in
	return
}

func (in *ProductSpec) DeepCopy() *ProductSpec {
	if in == nil {
		return nil
	}
	out := new(ProductSpec)
	in.DeepCopyInto(out)
	return out
}
